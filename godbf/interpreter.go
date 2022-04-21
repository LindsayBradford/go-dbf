package godbf

import (
	"bytes"
	"errors"
	"fmt"
)

// NewFromByteArray creates a DbfTable, reading it from a raw byte array, expecting the supplied encoding.
func NewFromByteArray(data []byte, fileEncoding string) (table *DbfTable, newErr error) {
	defer func() {
		if e := recover(); e != nil {
			newErr = fmt.Errorf("%v", e)
		}
	}()

	dt := new(DbfTable)
	dt.UseEncoding(fileEncoding)
	unpackHeader(data, dt)
	unpackRecords(data, dt)
	unpackFooter(data, dt)

	verifyTableAgainstRawBytes(data, dt)

	lockSchema(dt)
	return dt, nil
}

func unpackHeader(s []byte, dt *DbfTable) error {
	dt.fileSignature = s[0]
	dt.SetLastUpdatedFromBytes(s[1:4])
	dt.numberOfRecords = uint32(s[4]) | (uint32(s[5]) << 8) | (uint32(s[6]) << 16) | (uint32(s[7]) << 24)
	dt.numberOfBytesInHeader = uint16(s[8]) | (uint16(s[9]) << 8)
	dt.lengthOfEachRecord = uint16(s[10]) | (uint16(s[11]) << 8)

	if fieldErr := unpackFields(s, dt); fieldErr != nil {
		return fieldErr
	}
	return nil
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

func deriveFieldName(s []byte, dt *DbfTable, offset int) string {
	nameBytes := s[offset : offset+fieldNameByteLength]

	// Max usable field length is 10 bytes, where the 11th should contain the eod of field marker.
	endOfFieldIndex := bytes.Index(nameBytes, []byte{endOfFieldNameMarker})
	if endOfFieldIndex == -1 {
		msg := fmt.Sprintf("end-of-field marker missing from field bytes, offset [%d,%d]", offset, offset+fieldNameByteLength)
		panic(errors.New(msg))
	}

	fieldName := dt.decoder.ConvertString(string(nameBytes[:endOfFieldIndex]))
	return fieldName
}

func unpackRecords(data []byte, dt *DbfTable) {
	dt.dataStore = data // TODO: Deprecate?  At least reduce scope to just its records.
}

func unpackFooter(data []byte, dt *DbfTable) {
	dt.eofMarker = data[len(data)-1]
}

func verifyTableAgainstRawBytes(s []byte, dt *DbfTable) {
	verifyTableAgainstRawHeader(s, dt)
	verifyTableAgainstRawFooter(s, dt)
}

func verifyTableAgainstRawFooter(s []byte, dt *DbfTable) {
	if dt.eofMarker != eofMarker {
		panic(fmt.Errorf("encoded footer is %v, but actual footer is %d", eofMarker, s[len(s)-1]))
	}
}

func verifyTableAgainstRawHeader(s []byte, dt *DbfTable) {
	expectedSize := uint32(dt.numberOfBytesInHeader) + dt.numberOfRecords*uint32(dt.lengthOfEachRecord) + 1
	actualSize := uint32(len(s))
	if actualSize != expectedSize {
		panic(fmt.Errorf("encoded content is %d bytes, but header expected %d", actualSize, expectedSize))
	}
}

func lockSchema(dt *DbfTable) {
	dt.schemaLocked = true // Schema changes no longer permitted
}

// New creates a new dbase table from scratch for the given character encoding
func New(encoding string) (table *DbfTable) {

	// Create and populate DbaseTable struct
	dt := new(DbfTable)

	dt.UseEncoding(encoding)

	// set whether or not this table has been created from scratch
	dt.createdFromScratch = true

	// read dbase table header information
	dt.fileSignature = 0x03
	dt.RefreshLastUpdated()
	dt.numberOfRecords = 0
	dt.numberOfBytesInHeader = 32
	dt.lengthOfEachRecord = 0
	dt.fieldTerminator = 0x0D

	// create fieldMap to translate field name to index
	dt.fieldMap = make(map[string]int)

	// Number of fields in dbase table
	dt.numberOfFields = int((dt.numberOfBytesInHeader - 1 - 32) / 32)
	dt.eofMarker = eofMarker

	s := make([]byte, dt.numberOfBytesInHeader+1) // +1 is for footer

	//fmt.Printf("number of fields:\n%#v\n", numberOfFields)
	//fmt.Printf("DbfReader:\n%#v\n", int(dt.Fields[2].fixedFieldLength))

	//fmt.Printf("num records in table:%v\n", (dt.numberOfRecords))
	//fmt.Printf("fixedFieldLength of each record:%v\n", (dt.lengthOfEachRecord))

	// Since we are reading dbase file from the disk at least at this
	// phase changing schema of dbase file is not allowed.
	dt.schemaLocked = false

	// set DbfTable dataStore slice that will store the complete file in memory
	dt.dataStore = s

	dt.dataStore[0] = dt.fileSignature
	dt.dataStore[1] = dt.updateYear
	dt.dataStore[2] = dt.updateMonth
	dt.dataStore[3] = dt.updateDay

	// no MDX file (index upon demand)
	dt.dataStore[28] = 0x00

	// set dbase language driver
	// Huston we have problem!
	// There is no easy way to deal with encoding issues. At least at the moment
	// I will try to find archaic encoding code defined by dbase standard (if there is any)
	// for given encoding. If none match I will go with default ANSI.
	//
	// Despite this flag in set in dbase file, I will continue to use provide encoding for
	// the everything except this file encoding flag.
	//
	// Why? To make sure at least if you know the real encoding you can process text accordingly.

	if code, ok := encodingTable[lookup[encoding]]; ok {
		dt.dataStore[29] = code
	} else {
		dt.dataStore[29] = 0x57 // ANSI
	}

	dt.updateHeader()
	// no records as yet
	dt.dataStore = append(dt.dataStore, dt.eofMarker)

	return dt
}
