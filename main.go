package main

import (
	"benc2proto/proto-structs"
	"bufio"
	"bytes"
	"fmt"
)

type BencListElement interface {
	int32 | string
}

func decodeBencString(reader bufio.Reader, buffer bytes.Buffer) string {
	// TODO: Read and populate buffer till you find a :
	// TODO: Read buffer data and transform into int32
	// TODO: Read that many runes off reader
	// TODO: Cast to string and return
	// TODO: Reset buffer
	return ""
}

func decodeBencInt(reader bufio.Reader, buffer bytes.Buffer) int32 {
	// TODO: Empty out the buffer
	// TODO: Continue reading on the reader till you find an 'e'
	// TODO: Add each rune to buffer
	// TODO: Transform buffer data to int32
	// TODO: Reset buffer
	// TODO: return int32
	return 0
}

// TODO: Figure out how to have a list with 2 potential type elements, one option is custom struct which will provide both as separate properties
func decodeBencList(reader bufio.Reader, buffer bytes.Buffer) {
	// TODO: Empty out buffer
	// TODO: Read data from reader and add to buffer
	// TODO: Once we found a new property, pass to proper decoder
	// TODO: Add property/value to our list
	// TODO: Once we hit an 'e', return the list of data
}

func decodeBencDictionary(reader bufio.Reader, buffer bytes.Buffer) {
	// TODO: Empty out buffer
	// TODO: Read data from reader and add to buffer
	// TODO: Once we found a new property, pass to proper decoder
	// TODO: Once we have 2 pairs, add to dictionary
	// TODO: Once we hit an 'e', return the list of data
}

func DecodeTorrentFile(filename string) proto_structs.MetaInfo {
	// TODO: Open file reader
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
	fmt.Println(mi.String())
}
