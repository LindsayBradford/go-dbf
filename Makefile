include $(GOROOT)/src/Make.inc

TARG=godbf

GOFILES=src/pkg/godbf/dbfreader.go \
	src/pkg/godbf/helpers.go \
	
include $(GOROOT)/src/Make.pkg