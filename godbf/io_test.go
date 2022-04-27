package godbf

import (
	"errors"
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

const validTestFile = "testdata/validFile.dbf"
const lessThanActualRecordsFile = "testdata/lessThanActualRecords.dbf"

const realFile = "testdata/122016B1.DBF"

// For reference: https://en.wikipedia.org/wiki/.dbf#File_format_of_Level_5_DOS_dBASE

func TestNewFromFile_ValidFile_NoError(t *testing.T) {
	g := NewGomegaWithT(t)

	_, readError := NewFromFile(validTestFile, testEncoding)

	g.Expect(readError).To(BeNil())
}

func TestNewFromFile_ValidFile_TableIsCorrect(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest, _ := NewFromFile(validTestFile, testEncoding)

	t.Logf("DbfReader:\n%#v\n", tableUnderTest)
	t.Logf("tableUnderTest.FieldNames() = %v\n", tableUnderTest.FieldNames())
	t.Logf("tableUnderTest.NumberOfRecords() = %v\n", tableUnderTest.NumberOfRecords())
	t.Logf("tableUnderTest.lengthOfEachRecord  %v\n", tableUnderTest.lengthOfEachRecord)

	verifyTableIsCorrect(tableUnderTest, g)
}

func TestNewFromByteArray_TableIsCorrect(t *testing.T) {
	g := NewGomegaWithT(t)

	rawFileBytes, loadErr := ioutil.ReadFile(validTestFile)
	g.Expect(loadErr).To(BeNil())

	tableUnderTest, byteArrayErr := NewFromByteArray(rawFileBytes, testEncoding)
	g.Expect(byteArrayErr).To(BeNil())

	verifyTableIsCorrect(tableUnderTest, g)
}

func TestSaveToFile_LoadOfSavedIsCorrect(t *testing.T) {
	g := NewGomegaWithT(t)

	rawFileBytes, loadErr := ioutil.ReadFile(validTestFile)
	g.Expect(loadErr).To(BeNil())

	tableFromBytes, _ := NewFromByteArray(rawFileBytes, testEncoding)
	rawFileBytes = nil

	tempFilename := filepath.Join("testdata", "tempSavedTable.dbf")
	saveErr := SaveToFile(tableFromBytes, tempFilename)
	g.Expect(saveErr).To(BeNil())

	tableUnderTest, loadErr := NewFromFile(tempFilename, testEncoding)

	g.Expect(loadErr).To(BeNil())

	verifyTableIsCorrect(tableUnderTest, g)

	removeErr := os.Remove(tempFilename)
	g.Expect(removeErr).To(BeNil())
}

func TestSaveToFile_FromNew_NoError(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	sampleTime := tableUnderTest.LowDefTime(time.Now())

	g.Expect(tableUnderTest.NumberOfRecords()).To(BeZero())
	g.Expect(len(tableUnderTest.Fields())).To(BeZero())
	g.Expect(tableUnderTest.LastUpdated()).To(Equal(sampleTime))

	tempFilename := filepath.Join("testdata", "tempSavedTable.dbf")
	saveErr := SaveToFile(tableUnderTest, tempFilename)
	g.Expect(saveErr).To(BeNil())

	loadedTable, loadErr := NewFromFile(tempFilename, testEncoding)
	g.Expect(loadErr).To(BeNil())

	g.Expect(loadedTable.NumberOfRecords()).To(BeZero())
	g.Expect(len(loadedTable.Fields())).To(BeZero())
	g.Expect(loadedTable.LastUpdated()).To(Equal(sampleTime))

	removeErr := os.Remove(tempFilename)
	g.Expect(removeErr).To(BeNil())
}

func TestNewFromByteArray_EndOfFieldMarkerMissing_TableParsingError(t *testing.T) {
	g := NewGomegaWithT(t)

	rawFileBytes, loadErr := ioutil.ReadFile(validTestFile)
	g.Expect(loadErr).To(BeNil())

	// Pad entire name byte range, including the final 11th byte, with non-terminating characters.
	const startByteOfFirstFieldName = 32
	for i := startByteOfFirstFieldName; i <= startByteOfFirstFieldName+fieldNameByteLength; i++ {
		rawFileBytes[i] = 0x41 // UTF-8 'A'
	}

	_, byteArrayErr := NewFromByteArray(rawFileBytes, testEncoding)
	t.Log(byteArrayErr)

	g.Expect(byteArrayErr.Error()).To(ContainSubstring("end-of-field marker missing"))
}

func TestNewFromFile_NewFromLessThanActualRecords_Errors(t *testing.T) {
	g := NewGomegaWithT(t)

	_, readError := NewFromFile(lessThanActualRecordsFile, testEncoding)

	g.Expect(readError).ToNot(BeNil())
	t.Log(readError)
}

func verifyTableIsCorrect(tableUnderTest *DbfTable, g *GomegaWithT) {
	verifyFieldDescriptorsAreCorrect(tableUnderTest, g)
	verifyRecordsAreCorrect(tableUnderTest, g)
}

func verifyFieldDescriptorsAreCorrect(tableUnderTest *DbfTable, g *GomegaWithT) {
	expectedFieldNumber := 5
	fields := tableUnderTest.Fields()
	g.Expect(len(fields)).To(BeNumerically("==", expectedFieldNumber))

	expectedFieldNames := []string{"TESTBOOL", "TESTTEXT", "TESTDATE", "TESTNUM", "TESTFLOAT"}
	g.Expect(tableUnderTest.FieldNames()).To(Equal(expectedFieldNames))

	boolField := tableUnderTest.Fields()[0]
	g.Expect(boolField.fieldType).To(Equal(Logical))
	g.Expect(boolField.length).To(BeNumerically("==", 1))

	textField := tableUnderTest.Fields()[1]
	g.Expect(textField.fieldType).To(Equal(Character))
	g.Expect(textField.length).To(BeNumerically("==", 10))

	dateField := tableUnderTest.Fields()[2]
	g.Expect(dateField.fieldType).To(Equal(Date))
	g.Expect(dateField.length).To(BeNumerically("==", 8))

	numField := tableUnderTest.Fields()[3]
	g.Expect(numField.fieldType).To(Equal(Numeric))
	g.Expect(numField.length).To(BeNumerically("==", 10))
	g.Expect(numField.decimalPlaces).To(BeNumerically("==", 0))

	floatField := tableUnderTest.Fields()[4]
	g.Expect(floatField.fieldType).To(Equal(Float))
	g.Expect(floatField.length).To(BeNumerically("==", 10))
	g.Expect(floatField.decimalPlaces).To(BeNumerically("==", 2))
}

func verifyRecordsAreCorrect(tableUnderTest *DbfTable, g *GomegaWithT) {
	expectedRecordNumber := 3
	actualRecordNumber := tableUnderTest.NumberOfRecords()
	g.Expect(actualRecordNumber).To(BeNumerically("==", expectedRecordNumber))

	expectedRecord0Data := []string{"T", "test0", "20180101", "42", "42.01000"}
	g.Expect(tableUnderTest.GetRowAsSlice(0)).To(Equal(expectedRecord0Data))

	expectedRecord1Data := []string{"F", "test1", "20180102", "43", "43.02000"}
	g.Expect(tableUnderTest.GetRowAsSlice(1)).To(Equal(expectedRecord1Data))

	expectedRecord2Data := []string{"T", "test2", "20180103", "44", "44.03000"}
	g.Expect(tableUnderTest.GetRowAsSlice(2)).To(Equal(expectedRecord2Data))
}

func TestFieldsNameCorrectDetect(t *testing.T) {
	g := NewGomegaWithT(t)
	tableUnderTest, _ := NewFromFile(realFile, realEncoding)
	expectedFieldNumber := 18
	fields := tableUnderTest.Fields()
	g.Expect(len(fields)).To(BeNumerically("==", expectedFieldNumber))

	expectedFieldNames := []string{"REGN", "PLAN", "NUM_SC", "A_P", "VR", "VV", "VITG", "ORA", "OVA", "OITGA", "ORP", "OVP", "OITGP", "IR", "IV", "IITG", "DT", "PRIZ"}
	g.Expect(tableUnderTest.FieldNames()).To(Equal(expectedFieldNames))
}

func TestNewFromFile_ReaderPanics_Errors(t *testing.T) {
	g := NewGomegaWithT(t)

	reader = panicReader
	_, readError := NewFromFile(lessThanActualRecordsFile, testEncoding)

	g.Expect(readError).ToNot(BeNil())
	t.Log(readError)
}

func panicReader(r io.Reader, buf []byte) (int, error) {
	panic("I'm a little panic teapot")
}

func TestNewFromFile_ReaderErrors_Errors(t *testing.T) {
	g := NewGomegaWithT(t)

	reader = errorReader
	_, readError := NewFromFile(lessThanActualRecordsFile, testEncoding)

	g.Expect(readError).ToNot(BeNil())

	t.Log(readError)
}

func errorReader(r io.Reader, buf []byte) (int, error) {
	return -1, errors.New("i'm a little error teapot")
}

func TestNewFromFile_OpenErrors_Errors(t *testing.T) {
	g := NewGomegaWithT(t)

	fsWrapper = openErrorFileSystem{}
	_, readError := NewFromFile(lessThanActualRecordsFile, testEncoding)

	g.Expect(readError).ToNot(BeNil())
	t.Log(readError)
}

type openErrorFileSystem struct {
	osFileSystem
}

func (openErrorFileSystem) Open(name string) (file, error) {
	return nil, errors.New("Open error")
}

func TestNewFromFile_StatErrors_Errors(t *testing.T) {
	g := NewGomegaWithT(t)

	fsWrapper = statErrorFileSystem{}
	_, readError := NewFromFile(lessThanActualRecordsFile, testEncoding)

	g.Expect(readError).ToNot(BeNil())
	t.Log(readError)
}

type statErrorFileSystem struct {
	osFileSystem
}

func (statErrorFileSystem) Stat(name string) (os.FileInfo, error) {
	return nil, errors.New("Stat error")
}

func TestSaveToFile_CreateErrors_Errors(t *testing.T) {
	g := NewGomegaWithT(t)

	rawFileBytes, loadErr := ioutil.ReadFile(validTestFile)
	g.Expect(loadErr).To(BeNil())

	tableFromBytes, _ := NewFromByteArray(rawFileBytes, testEncoding)
	rawFileBytes = nil

	tempFilename := filepath.Join("testdata", "tempSavedTable.dbf")

	fsWrapper = createErrorFileSystem{}
	saveErr := SaveToFile(tableFromBytes, tempFilename)

	g.Expect(saveErr).ToNot(BeNil())
	t.Log(saveErr)
}

type createErrorFileSystem struct {
	osFileSystem
}

func (createErrorFileSystem) Create(name string) (*os.File, error) {
	return nil, errors.New("Create error")
}

func TestSaveToFile_CreatePanics_Errors(t *testing.T) {
	g := NewGomegaWithT(t)

	rawFileBytes, loadErr := ioutil.ReadFile(validTestFile)
	g.Expect(loadErr).To(BeNil())

	tableFromBytes, _ := NewFromByteArray(rawFileBytes, testEncoding)
	rawFileBytes = nil

	tempFilename := filepath.Join("testdata", "tempSavedTable.dbf")

	fsWrapper = createPanicFileSystem{}
	saveErr := SaveToFile(tableFromBytes, tempFilename)

	g.Expect(saveErr).ToNot(BeNil())
	t.Log(saveErr)
}

type createPanicFileSystem struct {
	osFileSystem
}

func (createPanicFileSystem) Create(name string) (*os.File, error) {
	panic(errors.New("Create panic"))
}
