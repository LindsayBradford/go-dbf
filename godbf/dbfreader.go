package godbf

import (
	"fmt"
	"errors"
	"os"
	"time"
	"code.google.com/p/mahonia"
	"strings"
	"strconv"
)

type DbfTable struct {
	// dbase file header information
	fileSignature         uint8 // Valid dBASE III PLUS table file (03h without a memo .DBT file; 83h with a memo)
	updateYear            uint8 // Date of last update; in YYMMDD format.
	updateMonth           uint8
	updateDay             uint8
	numberOfRecords       uint32   // Number of records in the table. 
	numberOfBytesInHeader uint16   // Number of bytes in the header. 
	lengthOfEachRecord    uint16   // Number of bytes in the record. 
	reservedBytes         [20]byte // Reserved bytes
	fieldDescriptor       [32]byte // Field descriptor array
	fieldTerminator       int8     // 0Dh stored as the field terminator.

	numberOfFields int // number of fiels/colums in dbase file

	// columns of dbase file
	fields []DbfField

	// used to map field names to index
	fieldMap map[string]int
	/*
		"dataEntryStarted" flag is used to control whether we can change
		dbase table structure when data enty started you can not change 
		the schema of the file if you are reading from an existing file this
		file will be set to "true". This means you can not modify the schema
		of a dbase table that you loaded from a file.
	*/
	dataEntryStarted bool

	// cratedFromScratch is used before adding new fields to increment nu
	cratedFromScratch bool

	// encoding of dbase file
	fileEncoding string
	// keeps the dbase table in memory as byte array
	dataStore []byte
}

type DbfField struct {
	fieldName   string
	fieldType   string
	fieldLength uint8
	fieldStore  [32]byte
}

func (df *DbfField) SetFieldName(fieldName string) {
	df.fieldName = fieldName
}

func NewFromFile(fileName string, fileEncoding string) (table *DbfTable, err error) {
	// create a decoder to decode file correctly
	d := mahonia.NewDecoder(fileEncoding)

	s, err := readFile(fileName)

	if err != nil {
		return nil, err
	}

	// Create and pupulate DbaseTable struct
	dt := new(DbfTable)

	dt.fileEncoding = fileEncoding

	// read dbase table header information
	dt.fileSignature = s[0]
	dt.updateYear = s[1]
	dt.updateMonth = s[2]
	dt.updateDay = s[3]
	dt.numberOfRecords = uint32(s[4]) | (uint32(s[5]) << 8) | (uint32(s[6]) << 16) | (uint32(s[7]) << 24)
	dt.numberOfBytesInHeader = uint16(s[8]) | (uint16(s[9]) << 8)
	dt.lengthOfEachRecord = uint16(s[10]) | (uint16(s[11]) << 8)

	// create fieldMap to taranslate field name to index
	dt.fieldMap = make(map[string]int)

	// Number of fields in dbase table
	dt.numberOfFields = int((dt.numberOfBytesInHeader - 1 - 32) / 32)

	// populate dbf fields
	for i := 0; i < int(dt.numberOfFields); i++ {
		offset := (i * 32) + 32

		fieldName := strings.Trim(d.ConvertString(string(s[offset:offset+10])), string([]byte{0}))
		dt.fieldMap[fieldName] = i

		var err error

		switch s[offset+11] {
		case 'C':
			err = dt.AddTextField(fieldName, s[offset+16])
		case 'N':
			err = dt.AddNumberField(fieldName, s[offset+16])
		case 'F':
			err = dt.AddFloatField(fieldName, s[offset+16])
		case 'L':
			err = dt.AddBooleanField(fieldName)
		case 'D':
			err = dt.AddDateField(fieldName)
		}

		// Check return value for errors
		if err != nil {
			return nil, err
		}

		//fmt.Printf("Field name:%v\n", fieldName)
		//fmt.Printf("Field data type:%v\n", string(s[offset+11]))
		//fmt.Printf("Field length:%v\n", s[offset+16])
		//fmt.Println("-----------------------------------------------")
	}

	//fmt.Printf("DbfReader:\n%#v\n", dt)
	//fmt.Printf("DbfReader:\n%#v\n", int(dt.Fields[2].fieldLength))

	//fmt.Printf("num records in table:%v\n", (dt.numberOfRecords))
	//fmt.Printf("lenght of each record:%v\n", (dt.lengthOfEachRecord))

	// Since we are reading dbase file from the disk at least at this
	// phase changing schema of dbase file is not allowed.
	dt.dataEntryStarted = true

	// set DbfTable dataStore slice that will store the complete file in memory
	dt.dataStore = s

	return dt, nil
}

// Create a new dbase table from the scratch
func New(encoding string) (table *DbfTable) {
	// create a decoder to decode file correctly
	//d := mahonia.NewDecoder(encoding)

	// Create and pupulate DbaseTable struct
	dt := new(DbfTable)

	dt.fileEncoding = encoding

	// set whether or not this table created from the scratch
	dt.cratedFromScratch = true
	
	// read dbase table header information
	dt.fileSignature = 0x03
	dt.updateYear = byte(time.Now().Year() % 100)
	dt.updateMonth = byte(time.Now().Month())
	dt.updateDay = byte(time.Now().YearDay())
	dt.numberOfRecords = 0
	dt.numberOfBytesInHeader = 32
	dt.lengthOfEachRecord = 0

	// create fieldMap to taranslate field name to index
	dt.fieldMap = make(map[string]int)

	// Number of fields in dbase table
	dt.numberOfFields = int((dt.numberOfBytesInHeader - 1 - 32) / 32)

	s := make([]byte, dt.numberOfBytesInHeader)

	//fmt.Printf("number of fields:\n%#v\n", numberOfFields)
	//fmt.Printf("DbfReader:\n%#v\n", int(dt.Fields[2].fieldLength))

	//fmt.Printf("num records in table:%v\n", (dt.numberOfRecords))
	//fmt.Printf("lenght of each record:%v\n", (dt.lengthOfEachRecord))

	// Since we are reading dbase file from the disk at least at this
	// phase changing schema of dbase file is not allowed.
	dt.dataEntryStarted = false

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
	// for given encoding. If none math I will go with default ANSI.
	//
	// Despite this flag in set in dbase file, I will continue to use provide encoding for 
	// the everything except this file encoding flag.
	//
	// Why? To make sure at least if you know the real encoding you can process text accordingly.
	
	if code, ok := encodingTable[lookup[encoding]]; ok {
		dt.dataStore[28] = code
	} else {
		dt.dataStore[28] = 0x57 // ANSI
	} 
	
	return dt
}

// Sets field value by index  
func (dt *DbfTable) SetFieldValueByName(row int, fieldName string, value string) (err error) {

	fieldIndex, ok := dt.fieldMap[fieldName]

	if !ok {
		return errors.New("Field name \"" + fieldName + "\" does not exist")
	}

	// set field value and return
	dt.SetFieldValue(row, fieldIndex, value)
	return
}

// Sets field value by name  
func (dt *DbfTable) SetFieldValue(row int, fieldIndex int, value string) (err error) {

	e := mahonia.NewEncoder(dt.fileEncoding)
	b := []byte(e.ConvertString(value))

	fieldLength := int(dt.fields[fieldIndex].fieldLength)

	//DEBUG

	//fmt.Printf("dt.numberOfBytesInHeader=%v\n\n", dt.numberOfBytesInHeader)
	//fmt.Printf("dt.lengthOfEachRecord=%v\n\n", dt.lengthOfEachRecord)

	// locate the offset of the field in DbfTable dataStore
	offset := int(dt.numberOfBytesInHeader)
	lengthOfRecord := int(dt.lengthOfEachRecord)

	offset = offset + (row * lengthOfRecord)

	recordOffset := 1

	for i := 0; i < len(dt.fields); i++ {
		if i == fieldIndex {
			break
		} else {
			recordOffset += int(dt.fields[i].fieldLength)
		}
	}

	// first fill the field with space values
	for i := 0; i < fieldLength; i++ {
		dt.dataStore[offset+recordOffset+i] = 0x20
	}

	// write new value
	switch dt.fields[fieldIndex].fieldType {
	case "C", "L", "D":
		for i := 0; i < len(b) && i < fieldLength; i++ {
			dt.dataStore[offset+recordOffset+i] = b[i]
		}
	case "N":
		for i := 0; i < fieldLength; i++ {
			// fmt.Printf("i:%v\n", i)
			if i < len(b) {
				dt.dataStore[offset+recordOffset+(fieldLength-i-1)] = b[(len(b)-1)-i]
			} else {
				break
			}
		}
	}

	return

	//fmt.Printf("field value:%#v\n", []byte(value))
	//fmt.Printf("field index:%#v\n", fieldIndex)
	//fmt.Printf("field length:%v\n", dt.Fields[fieldIndex].fieldLength)
	//fmt.Printf("string to byte:%#v\n", b)

}

func (dt *DbfTable) FieldValue(row int, fieldIndex int) (value string) {

	// create decoder to convert bytes to utf-8 	
	d := mahonia.NewDecoder(dt.fileEncoding)

	offset := int(dt.numberOfBytesInHeader)
	lengthOfRecord := int(dt.lengthOfEachRecord)

	offset = offset + (row * lengthOfRecord)

	recordOffset := 1

	for i := 0; i < len(dt.fields); i++ {
		if i == fieldIndex {
			break
		} else {
			recordOffset += int(dt.fields[i].fieldLength)
		}
	}

	temp := dt.dataStore[(offset + recordOffset):((offset + recordOffset) + int(dt.fields[fieldIndex].fieldLength))]
	
	for i:=0; i<len(temp); i++ {
		if temp[i] == 0x00 {
			temp = temp[0:i]
			break
		}
	}
	
	s := d.ConvertString(string(temp))
	//fmt.Printf("utf-8 value:[%#v]\n", s)
		
	value = strings.TrimSpace(s)
	
	
	//fmt.Printf("raw value:[%#v]\n", dt.dataStore[(offset + recordOffset):((offset + recordOffset) + int(dt.Fields[fieldIndex].fieldLength))])
	//fmt.Printf("utf-8 value:[%#v]\n", []byte(s))
	//value = string(dt.dataStore[(offset + recordOffset):((offset + recordOffset) + int(dt.Fields[fieldIndex].fieldLength))])
	return
}

// Float64FieldValueByName retuns the value of a field given row number and fieldName provided as a float64
func (dt *DbfTable) Float64FieldValueByName(row int, fieldName string) (value float64, err error) {

	fieldValueAsString, err := dt.FieldValueByName(row, fieldName)
	
	return strconv.ParseFloat(fieldValueAsString, 64)
}

// Int64FieldValueByName retuns the value of a field given row number and fieldName provided as an int64
func (dt *DbfTable) Int64FieldValueByName(row int, fieldName string) (value int64, err error) {

	fieldValueAsString, err := dt.FieldValueByName(row, fieldName)
	
	return strconv.ParseInt(fieldValueAsString, 0, 64)
}

// FieldValueByName retuns the value of a field given row number and fieldName provided
func (dt *DbfTable) FieldValueByName(row int, fieldName string) (value string, err error) {

	fieldIndex, ok := dt.fieldMap[fieldName]

	if !ok {
		err = errors.New("Field name \"" + fieldName + "\" does not exist")
		return
	}
	//fmt.Printf("fieldIndex:%v\n", fieldIndex)
	return dt.FieldValue(row, fieldIndex), err
}

func (dt *DbfTable) AddNewRecord() (newRecordNumber int) {

	if dt.dataEntryStarted == false {
		dt.dataEntryStarted = true
	}

	newRecord := make([]byte, dt.lengthOfEachRecord)
	dt.dataStore = appendSlice(dt.dataStore, newRecord)

	// since row numbers are "0" based first we set newRecordNumber
	// and then increment number of records in dbase table  
	newRecordNumber = int(dt.numberOfRecords)

	//fmt.Printf("Number of rows before:%d\n", dt.numberOfRecords)
	dt.numberOfRecords++
	s := uint32ToBytes(dt.numberOfRecords)
	dt.dataStore[4] = s[0]
	dt.dataStore[5] = s[1]
	dt.dataStore[6] = s[2]
	dt.dataStore[7] = s[3]
	//fmt.Printf("Number of rows after:%d\n", dt.numberOfRecords)

	return
}

func (dt *DbfTable) AddTextField(fieldName string, length uint8) (err error) {
	return dt.addField(fieldName, 'C', length)
}

func (dt *DbfTable) AddNumberField(fieldName string, length uint8) (err error) {
	return dt.addField(fieldName, 'N', length)
}

func (dt *DbfTable) AddFloatField(fieldName string, length uint8) (err error) {
	return dt.addField(fieldName, 'F', length)
}

func (dt *DbfTable) AddBooleanField(fieldName string) (err error) {
	return dt.addField(fieldName, 'L', 1)
}

func (dt *DbfTable) AddDateField(fieldName string) (err error) {
	return dt.addField(fieldName, 'D', 8)
}

// NumberOfRecords return number of rows in dbase table
func (dt *DbfTable) NumberOfRecords() int {
	return int(dt.numberOfRecords)
}

// Fields return slice of DbfField
func (dt *DbfTable) Fields() []DbfField {
	return dt.fields
}

// FieldName returns fieldName
func (df *DbfField) FieldName() string {
	return df.fieldName
}

// FieldType returns fieldType
func (df *DbfField) FieldType() string {
	return df.fieldType
}

// FieldName returns fieldName
func (df *DbfField) FieldLength() uint8 {
	return df.fieldLength
}

func (dt *DbfTable) addField(fieldName string, fieldType byte, length uint8) (err error) {

	if dt.dataEntryStarted {
		return errors.New("Once you start entering data to the dbase table or open an existing dbase file, altering dbase table schema is not allowed!")
	}

	s := dt.getNormalizedFieldName(fieldName)

	if dt.isFieldExist(s) {
		return errors.New("Field name \"" + s + "\" already exist!")
	}

	df := new(DbfField)
	df.fieldName = s
	df.fieldType = string(fieldType)
	df.fieldLength = length

	slice := dt.convertToByteSlice(s, 10)

	//fmt.Printf("len slice:%v\n", len(slice))

	// Field name in ASCII (max 10 chracters)
	for i := 0; i < len(slice); i++ {
		df.fieldStore[i] = slice[i]
		//fmt.Printf("i:%s\n", string(slice[i]))
	}

	// Field names are terminated by 00h	
	df.fieldStore[10] = 0x00

	// Set field's data type
	// C (Character) 	All OEM code page characters.
	// D (Date) 		Numbers and a character to separate month, day, and year (stored internally as 8 digits in YYYYMMDD format).
	// N (Numeric) 		- . 0 1 2 3 4 5 6 7 8 9
	// L (Logical) 		? Y y N n T t F f (? when not initialized). 
	df.fieldStore[11] = fieldType

	// length of field
	df.fieldStore[16] = length

	//fmt.Printf("addField | append:%v", df)

	dt.fields = append(dt.fields, *df)

	// if cratedFromScratch is true we need to update dbase header to reflect the changes we have made
	if dt.cratedFromScratch {
		dt.updateHeader()
	}

	return
}

// updateHeader updates the dbase file header after a field added
func (dt *DbfTable) updateHeader() {
	// first create a slice from initial 32 bytes of datastore as the foundation of the new slice
	// later we will set this slice to dt.dataStore to create the new header slice
	slice := dt.dataStore[0:32]

	// set dbase file signature
	slice[0] = 0x03

	var lengthOfEachRecord uint16 = 0

	for i := range dt.Fields() {
		lengthOfEachRecord += uint16(dt.Fields()[i].FieldLength())
		slice = appendSlice(slice, dt.Fields()[i].fieldStore[:])

		// don't forget to update fieldMap. We need it to find the index of a field name
		dt.fieldMap[dt.Fields()[i].FieldName()] = i
	}

	// end of file header terminator (0Dh)
	slice = appendSlice(slice, []byte{0x0D})

	// now reset dt.dataStore slice with the updated one
	dt.dataStore = slice

	// update the number of bytes in dbase file header
	dt.numberOfBytesInHeader = uint16(len(slice))
	s := uint32ToBytes(uint32(dt.numberOfBytesInHeader))
	dt.dataStore[8] = s[0]
	dt.dataStore[9] = s[1]

	dt.lengthOfEachRecord = lengthOfEachRecord + 1 // dont forget to add "1" for deletion marker which is 20h

	// update the lenght of each record
	s = uint32ToBytes(uint32(dt.lengthOfEachRecord))
	dt.dataStore[10] = s[0]
	dt.dataStore[11] = s[1]

	return
}

func (dt *DbfTable) SaveFile(filename string) (err error) {

	// don't forget to add dbase end of file marker which is 1Ah
	dt.dataStore = appendSlice(dt.dataStore, []byte{0x1A})

	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Hata:%v", err)
	}
	defer f.Close()

	n, err := f.Write(dt.dataStore)

	if err != nil {
		return err
	} else {
		fmt.Printf("%v bytes written to the file '%v'", n, filename)
	}

	return
}

func (dt *DbfTable) GetRowAsSlice(row int) []string {
	
	s := make([]string, len(dt.Fields()))
	
	for i := 0; i < len(dt.Fields()); i++ {
		s[i] = dt.FieldValue(row, i)
	}
	
	return s
}

func (dt *DbfTable) isFieldExist(fieldName string) bool {

	for i := 0; i < len(dt.fields); i++ {
		if dt.fields[i].fieldName == fieldName {
			return true
		}
	}

	return false
}

/*
	getByteSlice converts value to byte slice according to given encoding and return 
	a slice that is length equals to numberOfBytes or less if the string is shorter than
	numberOfBytes
*/
func (dt *DbfTable) convertToByteSlice(value string, numberOfBytes int) (s []byte) {
	e := mahonia.NewEncoder(dt.fileEncoding)
	b := []byte(e.ConvertString(value))

	if len(b) <= numberOfBytes {
		s = b
	} else {
		s = b[0:numberOfBytes]
	}
	return
}

func (dt *DbfTable) getNormalizedFieldName(name string) (s string) {
	e := mahonia.NewEncoder(dt.fileEncoding)
	b := []byte(e.ConvertString(name))

	if len(b) > 10 {
		b = b[0:10]
	}

	d := mahonia.NewDecoder(dt.fileEncoding)
	s = d.ConvertString(string(b))

	return
}
