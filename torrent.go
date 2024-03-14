package main

import (
	"benc2proto/proto_structs"
	"bufio"
	"crypto/sha1"
	"log"
	"os"

	"github.com/mitchellh/mapstructure"
	"google.golang.org/protobuf/proto"
)

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
