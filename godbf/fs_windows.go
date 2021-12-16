package godbf

import (
	"os"
	"syscall"
	"unsafe"
)

const (
	FILE_SHARE_READ   = 0x00000001
	FILE_SHARE_WRITE  = 0x00000002
	FILE_SHARE_DELETE = 0x00000004
)

var (
	shareMode uint32 = 0
)

func init() {
	if _, ok := os.LookupEnv("FILE_SHARE_READ"); ok {
		shareMode |= FILE_SHARE_READ
	}
	if _, ok := os.LookupEnv("FILE_SHARE_WRITE"); ok {
		shareMode |= FILE_SHARE_WRITE
	}
	if _, ok := os.LookupEnv("FILE_SHARE_DELETE"); ok {
		shareMode |= FILE_SHARE_DELETE
	}
}

func openFile(path string, mode int, perm uint32) (fd syscall.Handle, err error) {
	if len(path) == 0 {
		return syscall.InvalidHandle, syscall.ERROR_FILE_NOT_FOUND
	}
	pathp, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return syscall.InvalidHandle, err
	}
	var access uint32
	switch mode & (syscall.O_RDONLY | syscall.O_WRONLY | syscall.O_RDWR) {
	case syscall.O_RDONLY:
		access = syscall.GENERIC_READ
	case syscall.O_WRONLY:
		access = syscall.GENERIC_WRITE
	case syscall.O_RDWR:
		access = syscall.GENERIC_READ | syscall.GENERIC_WRITE
	}
	if mode&syscall.O_CREAT != 0 {
		access |= syscall.GENERIC_WRITE
	}
	if mode&syscall.O_APPEND != 0 {
		access &^= syscall.GENERIC_WRITE
		access |= syscall.FILE_APPEND_DATA
	}

	var sa *syscall.SecurityAttributes
	if mode&syscall.O_CLOEXEC == 0 {
		sa = makeInheritSa()
	}
	var createmode uint32
	switch {
	case mode&(syscall.O_CREAT|syscall.O_EXCL) == (syscall.O_CREAT | syscall.O_EXCL):
		createmode = syscall.CREATE_NEW
	case mode&(syscall.O_CREAT|syscall.O_TRUNC) == (syscall.O_CREAT | syscall.O_TRUNC):
		createmode = syscall.CREATE_ALWAYS
	case mode&syscall.O_CREAT == syscall.O_CREAT:
		createmode = syscall.OPEN_ALWAYS
	case mode&syscall.O_TRUNC == syscall.O_TRUNC:
		createmode = syscall.TRUNCATE_EXISTING
	default:
		createmode = syscall.OPEN_EXISTING
	}
	h, e := syscall.CreateFile(pathp, access, shareMode, sa, createmode, syscall.FILE_ATTRIBUTE_NORMAL, 0)
	return h, e
}

func makeInheritSa() *syscall.SecurityAttributes {
	var sa syscall.SecurityAttributes
	sa.Length = uint32(unsafe.Sizeof(sa))
	sa.InheritHandle = 1
	return &sa
}

func Open(name string) (file *os.File, err error) {
	return OpenFile(name, syscall.O_RDONLY, 0)
}

func OpenFile(name string, flag int, perm os.FileMode) (file *os.File, err error) {
	r, e := openFile(name, flag|syscall.O_CLOEXEC, syscallMode(perm))
	if e != nil {
		return nil, e
	}
	return os.NewFile(uintptr(r), name), nil
}

func syscallMode(i os.FileMode) (o uint32) {
	o |= uint32(i.Perm())
	if i&os.ModeSetuid != 0 {
		o |= syscall.S_ISUID
	}
	if i&os.ModeSetgid != 0 {
		o |= syscall.S_ISGID
	}
	if i&os.ModeSticky != 0 {
		o |= syscall.S_ISVTX
	}
	// No mapping for Go's ModeTemporary (plan9 only).
	return
}
