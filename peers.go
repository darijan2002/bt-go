package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"log"
	//"net"

	"fmt"
	"net/http"
	"net/url"

	"math/rand/v2"
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

func getPublicIp() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	defer resp.Body.Close()

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return string(ip), nil
}

func discoverPeers(filename string) ([]string, error) {

	hashInfo := HashInfo(DecodeTorrentFile(filename).Info)
	public, _ := getPublicIp()
	peer_id_buf := make([]byte, 20)
	binary.BigEndian.PutUint16(peer_id_buf, uint16(rand.IntN(50)))

	urlParams := url.Values{
		"info_hash":  []string{string(hashInfo[:])},
		"peer_addr":  []string{public},
		"downloaded": []string{"1000"},
		"uploaded":   []string{"100"},
		"peer_id":    []string{string(peer_id_buf)}, /* PEER ID : percent encoded of 20-byte array */
		"port":       []string{"17548"},
		"left":       []string{"100000"},
		"event":      []string{"started"},
		"compact":    []string{"0"},
	}

	target, err := url.Parse("http://35.204.74.244:7070/announce")

	if err != nil {
		return nil, err
	}

	target.RawQuery = urlParams.Encode()

	log.Println(target)
	resp, err := http.DefaultClient.Get(target.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	payload, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(bytes.NewReader(payload))
	d, _ := parseUpcomingProperty(r)

	peers := []string{}
	m := d.(map[string]any)
	log.Printf("Map data: %v", m)
	// TODO: Re-implement later on
	//ps := []byte(m["peers"].([]string))
	//for i := 0; i < len(ps); i += 6 {
	//	ip := net.IP(ps[i : i+4]).String()
	//	port := binary.BigEndian.Uint16([]byte(ps[i+4 : i+6]))
	//	peer := fmt.Sprintf("%s:%d", ip, port)
	//	peers = append(peers, peer)
	//}

	return peers, nil
}
