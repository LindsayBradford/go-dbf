package godbf

import (
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

func TestDbfTable_NewFromValidFile_FieldsCorrect(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest, _ := NewFromFile(validTestFile, testEncoding)

	expectedFieldNumber := 5
	fields := tableUnderTest.Fields()
	g.Expect(len(fields)).To(BeNumerically("==", expectedFieldNumber))

	expectedFieldNames := []string{"TESTBOOL", "TESTTEXT", "TESTDATE", "TESTNUM", "TESTFLOAT"}
	g.Expect(expectedFieldNames).To(Equal(tableUnderTest.FieldNames()))

	boolField := tableUnderTest.Fields()[0]
	g.Expect(boolField.fieldType).To(Equal(Logical))
	g.Expect(boolField.fieldLength).To(BeNumerically("==", 1))

	textField := tableUnderTest.Fields()[1]
	g.Expect(textField.fieldType).To(Equal(Character))
	g.Expect(textField.fieldLength).To(BeNumerically("==", 10))

	dateField := tableUnderTest.Fields()[2]
	g.Expect(dateField.fieldType).To(Equal(Date))
	g.Expect(dateField.fieldLength).To(BeNumerically("==", 8))

	numField := tableUnderTest.Fields()[3]
	g.Expect(numField.fieldType).To(Equal(Numeric))
	g.Expect(numField.fieldLength).To(BeNumerically("==", 10))
	g.Expect(numField.fieldDecimalPlaces).To(BeNumerically("==", 0))

	floatField := tableUnderTest.Fields()[4]
	g.Expect(floatField.fieldType).To(Equal(Float))
	g.Expect(floatField.fieldLength).To(BeNumerically("==", 10))
	g.Expect(floatField.fieldDecimalPlaces).To(BeNumerically("==", 2))
}

func TestDbfTable_NewFromValidFile_RecordsCorrect(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest, _ := NewFromFile(validTestFile, testEncoding)

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
