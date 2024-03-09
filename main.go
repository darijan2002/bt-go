package main

import (
	"benc2proto/proto_structs"
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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
		return decodeBencString(reader, buf), BencDataType(STRING)
	}

	switch rune {
	case 'i':
		return decodeBencInt(reader, buf), BencDataType(INT)
	case 'l':
		return decodeBencList(reader, buf), BencDataType(LIST)
	case 'd':
		return decodeBencDictionary(reader, buf), BencDataType(DICTIONARY)
	}

	return nil, BencDataType(UNKNOWN)
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
	for iSize := size; iSize > 0; iSize-- {
		rune, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				log.Println("Reached EOF")
				break
			}
			log.Printf("Errored during the reading of a rune %s %e\n", string(rune), err)
		}
		buffer.WriteRune(rune)
	}

	// Cast to string and return
	data, err := buffer.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Printf("Failed reading buffer, err is %e\n", err)
	}

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
	for true {
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
	for true {
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
		if keyType.String() != "STRING" {
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

func DecodeTorrentFile(filename string) proto_structs.MetaInfo {
	reader := OpenTorrentFile(filename)
	MapTorrentFile(reader) // TODO: Assign to variable
	// TODO: Map out torrentInfo data to MetaInfo
	return proto_structs.MetaInfo{}
}

func main() {
	mi := proto_structs.MetaInfo{
		Info: &proto_structs.Info{
			Name:        "name123",
			PieceLength: 1324,
			Pieces:      []string{"adf", "fsdfd"},
			Data: &proto_structs.Info_Files{
				Files: &proto_structs.FileInfos{Infos: []*proto_structs.FileInfo{
					{Length: 999, Path: "path/to/file"},
				}},
			},
		},
		Announce: "announce",
	}
	log.Println(mi.String())
}
