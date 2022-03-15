// Package godbf offers functionality for loading and saving  "dBASE Version 5" dbf formatted files.
// (https://en.wikipedia.org/wiki/.dbf#File_format_of_Level_5_DOS_dBASE) file structure.
// For the definitive source, see http://www.dbase.com/manuals/57LanguageReference.zip
package godbf

import (
	"bytes"
	"github.com/axgle/mahonia"
	"os"
)

// NewFromFile creates a DbfTable, reading it from a file with the given file name, expecting the supplied encoding.
func NewFromFile(fileName string, fileEncoding string) (*DbfTable, error) {
	data, readErr := readFile(fileName)
	if readErr != nil {
		return nil, readErr
	}
	return createDbfTable(data, fileEncoding)
}

// NewFromByteArray creates a DbfTable, reading it from a raw byte array, expecting the supplied encoding.
func NewFromByteArray(data []byte, fileEncoding string) (table *DbfTable, err error) {
	return createDbfTable(data, fileEncoding)
}

func createDbfTable(s []byte, fileEncoding string) (table *DbfTable, err error) {
	dt := new(DbfTable)
	assignEncoding(fileEncoding, dt)
	unpackHeader(s, dt)

	if fieldErr := unpackFields(s, dt); fieldErr != nil {
		return nil, fieldErr
	}

	finaliseSchema(s, dt)
	return dt, nil
}

func finaliseSchema(s []byte, dt *DbfTable) {
	dt.dataEntryStarted = true // Schema changes no longer permitted
	dt.dataStore = s           // TODO: Deprecate?
}

func unpackFields(s []byte, dt *DbfTable) error {
	// create fieldMap to translate field name to index
	dt.fieldMap = make(map[string]int)

	// Number of fields in dbase table
	dt.numberOfFields = int((dt.numberOfBytesInHeader - 1 - 32) / 32)
	for i := 0; i < dt.numberOfFields; i++ {
		if unpackFieldErr := unpackField(s, dt, i); unpackFieldErr != nil {
			return unpackFieldErr
		}
	}
	return nil
}

func unpackField(s []byte, dt *DbfTable, fieldIndex int) error {
	offset := (fieldIndex * 32) + 32
	fieldName := deriveFieldName(s, dt, offset)
	dt.fieldMap[fieldName] = fieldIndex

	var unpackErr error

	switch s[offset+11] {
	case 'C':
		unpackErr = dt.AddTextField(fieldName, s[offset+16])
	case 'N':
		unpackErr = dt.AddNumberField(fieldName, s[offset+16], s[offset+17])
	case 'F':
		unpackErr = dt.AddFloatField(fieldName, s[offset+16], s[offset+17])
	case 'L':
		unpackErr = dt.AddBooleanField(fieldName)
	case 'D':
		unpackErr = dt.AddDateField(fieldName)
	}

	if unpackErr != nil {
		return unpackErr
	}

	return nil
}

const endOfFieldMarker byte = 0x0

func deriveFieldName(s []byte, dt *DbfTable, offset int) string {
	nameBytes := s[offset : offset+maxFieldNameByteLength+1]
	// Max usable field length is 10 bytes, where the 11th is guaranteed to contain the eod of field marker.
	endOfFieldIndex := bytes.Index(nameBytes, []byte{endOfFieldMarker})
	fieldName := dt.encoder.ConvertString(string(nameBytes[:endOfFieldIndex]))
	return fieldName
}

func unpackHeader(s []byte, dt *DbfTable) {
	dt.fileSignature = s[0]

	dt.updateYear = s[1]
	dt.updateMonth = s[2]
	dt.updateDay = s[3]

	dt.numberOfRecords = uint32(s[4]) | (uint32(s[5]) << 8) | (uint32(s[6]) << 16) | (uint32(s[7]) << 24)
	dt.numberOfBytesInHeader = uint16(s[8]) | (uint16(s[9]) << 8)
	dt.lengthOfEachRecord = uint16(s[10]) | (uint16(s[11]) << 8)
}

func assignEncoding(fileEncoding string, dt *DbfTable) {
	dt.fileEncoding = fileEncoding
	dt.encoder = mahonia.NewEncoder(fileEncoding)
	dt.decoder = mahonia.NewDecoder(fileEncoding)
}

// SaveToFile saves the supplied DbfTable to a file of the specified filename
func SaveToFile(dt *DbfTable, filename string) (saveErr error) {
	f, createErr := os.Create(filename)
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
	if dsErr := writeDataStore(dt, f); dsErr != nil {
		return dsErr
	}
	if footerErr := writeFooter(f); footerErr != nil {
		return footerErr
	}
	return nil
}

func writeDataStore(dt *DbfTable, f *os.File) error {
	if _, dsErr := f.Write(dt.dataStore); dsErr != nil {
		return dsErr
	}
	return nil
}

const EofMarker byte = 0x1A

func writeFooter(f *os.File) error {
	eofBytes := []byte{EofMarker}
	if _, footerErr := f.Write(eofBytes); footerErr != nil {
		return footerErr
	}
	return nil
}
