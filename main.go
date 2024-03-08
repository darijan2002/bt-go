package benc2proto

import (
	proto_structs "benc2proto/proto-structs"
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type BencListElement interface {
	int32 | string
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
		log.Println("Error during conversion")
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

// TODO: Figure out how to have a list with 2 potential type elements,
// one option is custom struct which will provide both as separate properties
func decodeBencList(reader *bufio.Reader, buffer *bytes.Buffer) {
	// TODO: Empty out buffer
	buffer.Reset()

	// TODO: Read data from reader and add to buffer

	// TODO: Once we found a new property, pass to proper decoder

	// TODO: Add property/value to our list

	// TODO: Once we hit an 'e', return the list of data
}

func decodeBencDictionary(reader *bufio.Reader, buffer *bytes.Buffer) {
	// TODO: Empty out buffer
	buffer.Reset()

	// TODO: Read data from reader and add to buffer

	// TODO: Once we found a new property, pass to proper decoder

	// TODO: Once we have 2 pairs, add to dictionary

	// TODO: Once we hit an 'e', return the list of data
}

func parseUpcomingProperty(reader *bufio.Reader) (string, int, BencDataType) {
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	rune, _, err := reader.ReadRune()
	if err != nil {
		log.Println("Error while reading initial rune")
	}

	buf.WriteRune(rune)

	if _, err := strconv.Atoi(strconv.QuoteRune(rune)); err == nil {
		return decodeBencString(reader, buf), 0, BencDataType(INT)
	}

	switch rune {
	case 'i':
		return "", decodeBencInt(reader, buf), BencDataType(STRING)
		// TODO: Expand with list and dictionary
	}

	return "", 0, BencDataType(UNKNOWN)
}

func DecodeTorrentFile(filename string) proto_structs.MetaInfo {
	// TODO: Open file reader
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Error while reading file")
	}

	reader := bufio.NewReader(file)

	parseUpcomingProperty(reader)
	// TODO: Check line type

	// TODO: Process line properly (aka, if d<bencode>e, decode as map)

	// TODO: Repeat
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
