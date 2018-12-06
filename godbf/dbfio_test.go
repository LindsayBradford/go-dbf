package godbf

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
)

const validTestFile = "testdata/validFile.dbf"

// For reference: https://www.dbase.com/Knowledgebase/INT/db7_file_fmt.htm

func TestDbfTable_NewFromValidFile_NoError(t *testing.T) {
	g := NewGomegaWithT(t)

	_, readError := NewFromFile(validTestFile, testEncoding)

	g.Expect(readError).To(BeNil())
}

func TestDbfTable_NewFromValidFile_FieldDescriptorsCorrect(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest, _ := NewFromFile(validTestFile, testEncoding)

	verifyFieldDescriptorsAreCorrect(tableUnderTest, g)
}

func verifyFieldDescriptorsAreCorrect(tableUnderTest *DbfTable, g *GomegaWithT) {
	expectedFieldNumber := 5
	fields := tableUnderTest.Fields()
	g.Expect(len(fields)).To(BeNumerically("==", expectedFieldNumber))

	expectedFieldNames := []string{"TESTBOOL", "TESTTEXT", "TESTDATE", "TESTNUM", "TESTFLOAT"}
	g.Expect(expectedFieldNames).To(Equal(tableUnderTest.FieldNames()))

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

func TestDbfTable_NewFromValidFile_RecordsCorrect(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest, _ := NewFromFile(validTestFile, testEncoding)

	verifyRecordsAreCorrect(tableUnderTest, g)
}

func verifyRecordsAreCorrect(tableUnderTest *DbfTable, g *GomegaWithT) {
	expectedRecordNumber := 3
	actualRecordNumber := tableUnderTest.NumberOfRecords()
	g.Expect(expectedRecordNumber).To(BeNumerically("==", actualRecordNumber))

	expectedRecord0Data := []string{"T", "test0", "20180101", "42", "42.01000"}
	g.Expect(tableUnderTest.GetRowAsSlice(0)).To(Equal(expectedRecord0Data))

	expectedRecord1Data := []string{"F", "test1", "20180102", "43", "43.02000"}
	g.Expect(tableUnderTest.GetRowAsSlice(1)).To(Equal(expectedRecord1Data))

	expectedRecord2Data := []string{"T", "test2", "20180103", "44", "44.03000"}
	g.Expect(tableUnderTest.GetRowAsSlice(2)).To(Equal(expectedRecord2Data))
}

func TestDbfTable_SaveFile_LoadOfSavedIsCorrect(t *testing.T) {
	g := NewGomegaWithT(t)

	validTable, _ := NewFromFile(validTestFile, testEncoding)

	tempFile := filepath.Join("testdata", "tempSavedTable.dbf")

	validTable.SaveFile(tempFile)
	validTable = nil

	tableUnderTest, loadErr := NewFromFile(tempFile, testEncoding)

	g.Expect(loadErr).To(BeNil())

	verifyFieldDescriptorsAreCorrect(tableUnderTest, g)
	verifyRecordsAreCorrect(tableUnderTest, g)

	os.Remove(tempFile)
}

func TestDbfTable_NewFromByteArray_TableIsCorrect(t *testing.T) {
	g := NewGomegaWithT(t)

	validTable, _ := NewFromFile(validTestFile, testEncoding)

	tableUnderTest, byteArrayErr := NewFromByteArray(validTable.dataStore, testEncoding)

	g.Expect(byteArrayErr).To(BeNil())

	verifyFieldDescriptorsAreCorrect(tableUnderTest, g)
	verifyRecordsAreCorrect(tableUnderTest, g)
}
