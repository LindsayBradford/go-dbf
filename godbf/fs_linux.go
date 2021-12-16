package godbf

import (
	"os"
)

func Open(name string) (file *os.File, err error) {
	return os.Open(name)
}
