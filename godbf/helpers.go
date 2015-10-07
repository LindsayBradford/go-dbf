package godbf

import (
	"io"
	"os"
	//"fmt"
)

// append data[]byte to slice and returns the new byte slice
func appendSlice(slice, data []byte) []byte {
	l := len(slice)
	if l+len(data) > cap(slice) { // reallocate
		// Allocate double what's needed, for future growth.
		// increase capcity %20
		//cap := (l+len(data))*(5/4)
		//newSlice := make([]byte, cap)
		newSlice := make([]byte, (l+len(data))*2)
		// The copy function is predeclared and works for any slice type.
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	//fmt.Printf("appendSlice son:%v\n", len(slice))
	//fmt.Printf("appendSlice cap:%v\n", cap(slice))
	for i, c := range data {
		slice[l+i] = c
	}
	return slice
}

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

func int32ToBytes(x int32) []byte {
	var buf [4]byte
	buf[0] = byte(x >> 0)
	buf[1] = byte(x >> 8)
	buf[2] = byte(x >> 16)
	buf[3] = byte(x >> 24)
	return buf[:]
}

func uint64ToBytes(x uint64) []byte {
	var buf [8]byte
	buf[0] = byte(x >> 0)
	buf[1] = byte(x >> 8)
	buf[2] = byte(x >> 16)
	buf[3] = byte(x >> 24)
	buf[4] = byte(x >> 32)
	buf[5] = byte(x >> 40)
	buf[6] = byte(x >> 48)
	buf[7] = byte(x >> 56)

	return buf[:]
}

func int64ToBytes(x uint64) []byte {
	var buf [8]byte
	buf[0] = byte(x >> 0)
	buf[1] = byte(x >> 8)
	buf[2] = byte(x >> 16)
	buf[3] = byte(x >> 24)
	buf[4] = byte(x >> 32)
	buf[5] = byte(x >> 40)
	buf[6] = byte(x >> 48)
	buf[7] = byte(x >> 56)

	return buf[:]
}

func bytesToInt32le(b []byte) int32 {
	return int32(uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24)

}

func bytesToInt32be(b []byte) int32 {
	return int32(uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24)
}

func bytesToUint64le(b []byte) uint64 {
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}
