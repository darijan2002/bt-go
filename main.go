package main

import (
	"benc2proto/proto_structs"
	"bufio"
	"bytes"
	"crypto/sha1"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/mapstructure"
	"google.golang.org/protobuf/proto"
)

// NOTE: Try to figure out where we can use defer to optimize performance by closing the files and readers which aren't in use any more

func parseUpcomingProperty(reader *bufio.Reader) (any, BencDataType) {
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	rune, _, err := reader.ReadRune()
	if err != nil {
		log.Println("Error while reading initial rune")
	}

	buf.WriteRune(rune)

	if _, err := strconv.Atoi(string(rune)); err == nil {
		return decodeBencString(reader, buf), BENC_STRING
	}

	switch rune {
	case 'i':
		return decodeBencInt(reader, buf), BENC_INT
	case 'l':
		return decodeBencList(reader, buf), BENC_LIST
	case 'd':
		return decodeBencDictionary(reader, buf), BENC_DICTIONARY
	}

	return nil, BENC_UNKNOWN
}

func decodeBencString(reader *bufio.Reader, buffer *bytes.Buffer) string {
	// Read and populate buffer till you find a :
	sizeStr, err := reader.ReadString(':')

	// Read buffer data and transform into int32
	buffer.WriteString(strings.TrimSuffix(sizeStr, ":"))
	fs, err := buffer.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Printf("Failed reading buffer, err is %e\n", err)
	}
	size, err := strconv.Atoi(fs)
	if err != nil {
		log.Printf("Error during conversion: %e", err)
	}

	// Read that many runes off reader
	limReader := io.LimitReader(reader, int64(size))
	bytesRead, err := buffer.ReadFrom(limReader)

	if err != nil && err != io.EOF {
		log.Printf("Failed reading buffer, err is %e\n", err)
	}
	if bytesRead != int64(size) {
		log.Println("Read less bytes than needed!")
	}

	// Cast to string and return
	data := buffer.String()

	// Reset buffer
	buffer.Reset()
	return data
}

func decodeBencInt(reader *bufio.Reader, buffer *bytes.Buffer) int {
	// Empty out the buffer
	buffer.Reset()

	// Continue reading on the reader till you find an 'e'
	numberStr, err := reader.ReadString('e')
	if err != nil {
		log.Println("Failed to read integer")
	}

	// Convert number to int
	number, err := strconv.Atoi(strings.TrimSuffix(numberStr, "e"))
	if err != nil {
		log.Println("Failed to convert to integer")
	}

	return number
}

func decodeBencList(reader *bufio.Reader, buffer *bytes.Buffer) []any {
	// Empty out buffer
	buffer.Reset()

	data := []any{}
	for {
		rune, _, err := reader.ReadRune()

		// Once we hit an 'e', return the list of data
		if err != nil || rune == 'e' {
			break
		}

		// rewind reader so we don't corrupt the pointer position
		reader.UnreadRune()

		// Once we found a new property, pass to proper decoder
		item, _ := parseUpcomingProperty(reader)

		// Add property/value to our list
		data = append(data, item)
	}

	return data
}

func decodeBencDictionary(reader *bufio.Reader, buffer *bytes.Buffer) map[string]any {
	// Empty out buffer
	buffer.Reset()

	data := map[string]any{}
	// Read data from reader and add to buffer
	for {
		rune, _, err := reader.ReadRune()

		// Once we hit an 'e', return the list of data
		if err != nil || rune == 'e' {
			break
		}
		// Rewind reader pointer so we don't corrupt data reading flow when parsing
		reader.UnreadRune()

		// Parse the key and value
		key, keyType := parseUpcomingProperty(reader)
		val, _ := parseUpcomingProperty(reader)
		// Ensure we are reading a string
		if keyType != BENC_STRING {
			log.Println("Failed reading key, not a string")
		}

		// Cast key to string and assign in 'data'
		data[key.(string)] = val
	}
	return data
}

func OpenTorrentFile(filename string) *bufio.Reader {
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Error while reading file")
	}

	return bufio.NewReader(file)
}

func MapTorrentFile(reader *bufio.Reader) map[string]any {
	torrentInfo, _ := parseUpcomingProperty(reader)
	return torrentInfo.(map[string]any)
}

func DecodeTorrentFile(filename string) *proto_structs.MetaInfo {
	reader := OpenTorrentFile(filename)
	parsedTorrent := MapTorrentFile(reader)

	// Split the pieces into arrays of 20 bytes each
	pieces := []byte(parsedTorrent["info"].(map[string]any)["pieces"].(string))
	log.Println("'pieces' field has", len(pieces), "characters")
	n := len(pieces) / 20

	list := make([][20]byte, n)
	for i := 0; i < n; i++ {
		list[i] = [20]byte(pieces[20*i : 20*(i+1)])
	}
	parsedTorrent["info"].(map[string]any)["pieces"] = list

	res := &proto_structs.MetaInfo{}
	err := mapstructure.Decode(parsedTorrent, res)
	if err != nil {
		log.Println(err)
	}

	return res
}

func HashInfo(info *proto_structs.Info) [20]byte {
	// Encode info to protobuf
	info_encoded, err := proto.Marshal(info)
	if err != nil {
		log.Panicln("Failed to encode info struct to wire enc!")
	}

	// Hash with SHA1
	hash := sha1.Sum(info_encoded)
	return hash
}

func main() {
	//printSampleMI()

	metainf := DecodeTorrentFile("torrents/big-buck-bunny.torrent") // len(pieces) = 21100 -> 1055 pieces
	//metainf := DecodeTorrentFile("torrents/sample.torrent") // len(pieces) = 60 -> 60/20 = 3 pieces
	log.Println("Torrent has", len(metainf.Info.Pieces), "pieces")
	//fmt.Printf("%v %T\n", metainf, metainf)

	hash := HashInfo(metainf.Info)
	log.Println("Hash of info struct:", hash, len(hash))
}
