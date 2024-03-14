package main

import (
	"log"
)

// NOTE: Try to figure out where we can use defer to optimize performance by closing the files and readers which aren't in use any more

func main() {
	//printSampleMI()

	metainf := DecodeTorrentFile("torrents/big-buck-bunny.torrent") // len(pieces) = 21100 -> 1055 pieces
	//metainf := DecodeTorrentFile("torrents/sample.torrent") // len(pieces) = 60 -> 60/20 = 3 pieces
	log.Println("Torrent has", len(metainf.Info.Pieces), "pieces")
	//fmt.Printf("%v %T\n", metainf, metainf)

	peers, _ := discoverPeers("torrents/big-buck-bunny.torrent")
	log.Printf("Peers: %v", peers)

	for true {

	}

	hash := HashInfo(metainf.Info)
	log.Println("Hash of info struct:", hash, len(hash))
}
