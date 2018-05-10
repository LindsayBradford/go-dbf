package godbf

import (
	"fmt"
	"os"
	"strings"

	"github.com/axgle/mahonia"
)

func NewFromFile(fileName string, fileEncoding string) (table *DbfTable, err error) {
	s, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	return createDbfTable(nil, s, fileEncoding)
}

func NewFromByteArray(data []byte, fileEncoding string) (table *DbfTable, err error) {
	return createDbfTable(data, nil, fileEncoding)
}

func createDbfTable(data []byte, file *os.File, fileEncoding string) (table *DbfTable, err error) {
	// Create and pupulate DbaseTable struct
	dt := new(DbfTable)

	dt.fileEncoding = fileEncoding
	dt.dataFile = file
	dt.dataStore = data
	dt.useMemory = file == nil

	dt.encoder = mahonia.NewEncoder(fileEncoding)
	dt.decoder = mahonia.NewDecoder(fileEncoding)

	s := make([]byte, 12)

	if _, err := dt.readAt(s, 0); err != nil {
		return nil, err
	}

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

	s = make([]byte, int(dt.numberOfFields)*32+32)

	if _, err := dt.readAt(s, 0); err != nil {
		return nil, err
	}

	// populate dbf fields
	for i := 0; i < int(dt.numberOfFields); i++ {
		offset := (i * 32) + 32

		fieldName := strings.Trim(dt.encoder.ConvertString(string(s[offset:offset+10])), string([]byte{0}))
		dt.fieldMap[fieldName] = i

		var err error

		switch s[offset+11] {
		case 'C':
			err = dt.AddTextField(fieldName, s[offset+16])
		case 'N':
			err = dt.AddNumberField(fieldName, s[offset+16], s[offset+17])
		case 'F':
			err = dt.AddFloatField(fieldName, s[offset+16], s[offset+17])
		case 'L':
			err = dt.AddBooleanField(fieldName)
		case 'D':
			err = dt.AddDateField(fieldName)
		case 'I':
			err = dt.AddImageField(fieldName, s[offset+16], s[offset+17])
		default:
			println(`Found incorrect code field: `, string(s[offset+11]), `use as Text field: `, fieldName)
			err = dt.AddTextField(fieldName, s[offset+16])
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

	return dt, nil
}

func (dt *DbfTable) SaveFile(filename string) (err error) {

	f, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer f.Close()

	dsBytes, dsErr := f.Write(dt.dataStore)

	if dsErr != nil {
		return dsErr
	}

	// Add dbase end of file marker (1Ah)

	footerByte, footerErr := f.Write([]byte{0x1A})

	if footerErr != nil {
		return footerErr
	}

	fmt.Printf("%v bytes written to file '%v'.\n", dsBytes+footerByte, filename)

	return nil
}

func (dt *DbfTable) Close() error {
	if !dt.useMemory {
		return dt.dataFile.Close()
	}
	dt.dataStore = make([]byte, 0)
	return nil
}
