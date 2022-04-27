// Package godbf offers functionality for loading and saving  "dBASE Version 5" dbf formatted files.
// (https://en.wikipedia.org/wiki/.dbf#File_format_of_Level_5_DOS_dBASE) file structure.
// For the definitive source, see http://www.dbase.com/manuals/57LanguageReference.zip
package godbf

import (
	"fmt"
	"os"
)

// NewFromFile creates a DbfTable, reading it from a file with the given file name, expecting the supplied encoding.
func NewFromFile(fileName string, fileEncoding string) (table *DbfTable, newErr error) {
	defer func() {
		if e := recover(); e != nil {
			newErr = fmt.Errorf("%v", e)
		}
	}()

	data, readErr := readFile(fileName)
	if readErr != nil {
		return nil, readErr
	}
	return NewFromByteArray(data, fileEncoding)
}

// SaveToFile saves the supplied DbfTable to a file of the specified filename
func SaveToFile(dt *DbfTable, filename string) (saveErr error) {
	defer func() {
		if e := recover(); e != nil {
			saveErr = fmt.Errorf("%v", e)
		}
	}()

	f, createErr := fsWrapper.Create(filename)
	if createErr != nil {
		return createErr
	}

	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			saveErr = closeErr
		}
	}()

	writeErr := writeContent(dt, f)
	if writeErr != nil {
		return writeErr
	}

	return saveErr
}

func writeContent(dt *DbfTable, f *os.File) error {
	if _, dsErr := f.Write(dt.dataStore); dsErr != nil {
		return dsErr
	}
	return nil
}
