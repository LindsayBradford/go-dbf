package godbf

import (
	"io"
	"os"
)

func readFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	d, err := f.Stat()
	if err != nil {
		return nil, err
	}

	buf := make([]byte, d.Size())
	_, err = io.ReadFull(f, buf)
	if err != nil {
		return nil, err
	}
	return buf, err
}

func uint32ToBytes(x uint32) []byte {
	var buf [4]byte
	buf[0] = byte(x >> 0)
	buf[1] = byte(x >> 8)
	buf[2] = byte(x >> 16)
	buf[3] = byte(x >> 24)
	return buf[:]
}
