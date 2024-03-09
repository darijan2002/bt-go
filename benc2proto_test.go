package main

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDecodeToString(tc *testing.T) {
	reader := bufio.NewReader(strings.NewReader("0:hellohello"))
	buf := bytes.NewBufferString("1")
	data := decodeBencString(reader, buf)
	assert.Equal(tc, "hellohello", data, "String mismatched")
}

func TestDecodeToList(tc *testing.T) {
	reader := bufio.NewReader(strings.NewReader("10:hellohelloi32ee"))
	buf := bytes.NewBufferString("l")
	data := decodeBencList(reader, buf)
	assert.Equal(tc, []any{"hellohello", 32}, data, "Array is incorrect")
}

func TestDecodeToDictionary(tc *testing.T) {
	reader := bufio.NewReader(strings.NewReader("10:hellohelloi32ee"))
	buf := bytes.NewBufferString("d")
	data := decodeBencDictionary(reader, buf)
	assert.Equal(tc, map[string]any{"hellohello": 32}, data, "Map is incorrect")
}

func TestDecodeToDictionaryWithListValue(tc *testing.T) {
	reader := bufio.NewReader(strings.NewReader("10:hellohellol5:helloi32eee"))
	buf := bytes.NewBufferString("d")
	data := decodeBencDictionary(reader, buf)
	assert.Equal(tc, map[string]any{"hellohello": []any{"hello", 32}}, data, "Map is incorrect")
}

func TestOpenAndMapTorrentFile(tc *testing.T) {
	reader := OpenTorrentFile("./torrents/big-buck-bunny.torrent")
	torrentInfo := MapTorrentFile(reader)
	importantProps := []string{"info", "announce"}

	keys := make([]string, 0, len(torrentInfo))
	for key, _ := range torrentInfo {
		keys = append(keys, key)
	}

	for _, prop := range importantProps {
		assert.Contains(tc, keys, prop, "Important property field doesn't exist in torrent info")
	}
}
