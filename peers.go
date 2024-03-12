package peers

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
	"main"
)

// Tracker
/*
http://0.0.0.0:7070/announce?
	info_hash=%81%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00%00&
	peer_addr=2.137.87.41&
	downloaded=0&
	uploaded=0&
	peer_id=-qB00000000000000001&
	port=17548&
	left=0&
	event=completed&
	compact=0
*/

func getPublicIp() string {
	resp, err := http.Get("https://api.ipify.org?format=text")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    ip, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

	return string(ip)
}

func discoverPeers(filename string) []string {

	urlParams := url.Values{
		"info_hash" : string(main.HashInfo(main.DecodeTorrentFile(filename).Info)),
		"peer_addr" : getPublicIp,
		"downloaded" : "0",
		"uploaded" : "0",
		"peer_id" : /* PEER ID : percent encoded of 20-byte array */,
		"port" : "17548",
		"left" : "0",
		"event" : "completed",
		"compact" : "0"
	}

	target, err := url.Parse("http://0.0.0.0:7070/announce")

	if err != nil {
		return peers, err
	}

	target.RawQuery = urlParams.Encode()

	resp, err := http.DefaultClient.Get(target.String())
	if err != nil {
		return peers, err
	}
	defer resp.Body.Close()

	payload, err := io.ReadAll(resp.Body)

	if err != nil {
		return peers, err
	}

	r := bytes.NewReader(payload)
	d, err := decode.parseUpcomingProperty(r)

	if err != nil {
		return peers, err
	}

	m := d.(map[string]any)
	ps := []byte(m["peers"].(string))
	for i := 0; i < len(ps); i += 6 {
		ip := net.IP(ps[i : i+4]).String()
		port := binary.BigEndian.Uint16([]byte(ps[i+4 : i+6]))
		peer := fmt.Sprintf("%s:%d", ip, port)
		peers = append(peers, peer)
	}

	return peers, nil
}