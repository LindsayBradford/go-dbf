package godbf

import (
	"io"
	"os"
)

type readerFunction func(r io.Reader, buf []byte) (int, error)

var reader readerFunction = io.ReadFull

// https://talks.golang.org/2012/10things.slide#8

type fileSystem interface {
	Create(name string) (*os.File, error)
	Open(name string) (file, error)
	Stat(name string) (os.FileInfo, error)
}

// osFileSystem implements fileSystem using the local disk.
type osFileSystem struct{}

func (osFileSystem) Create(name string) (*os.File, error)  { return os.Create(name) }
func (osFileSystem) Open(name string) (file, error)        { return os.Open(name) }
func (osFileSystem) Stat(name string) (os.FileInfo, error) { return os.Stat(name) }

var fsWrapper fileSystem = osFileSystem{}

type file interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	Stat() (os.FileInfo, error)
}

func readFile(filename string) ([]byte, error) {
	f, err := fsWrapper.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	d, err := fsWrapper.Stat(filename)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, d.Size())
	_, err = reader(f, buf)
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
