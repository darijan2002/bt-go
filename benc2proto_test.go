package benc2proto

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
