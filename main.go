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

	"fmt"
    "io/ioutil"
    "net/http"
	"decode"
	"peers"
	"torrent"
)

// NOTE: Try to figure out where we can use defer to optimize performance by closing the files and readers which aren't in use any more

func main() {
	//printSampleMI()

	metainf := torrent.DecodeTorrentFile("torrents/big-buck-bunny.torrent") // len(pieces) = 21100 -> 1055 pieces
	//metainf := DecodeTorrentFile("torrents/sample.torrent") // len(pieces) = 60 -> 60/20 = 3 pieces
	log.Println("Torrent has", len(metainf.Info.Pieces), "pieces")
	//fmt.Printf("%v %T\n", metainf, metainf)

	hash := torrent.HashInfo(metainf.Info)
	log.Println("Hash of info struct:", hash, len(hash))
}
