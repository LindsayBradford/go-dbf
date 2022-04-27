package godbf

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

const testEncoding = "UTF-8"
const realEncoding = "cp866"

// For reference: https://en.wikipedia.org/wiki/.dbf#File_format_of_Level_5_DOS_dBASE

func TestDbfTable_New(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	g.Expect(tableUnderTest.NumberOfRecords()).To(BeZero())
	g.Expect(len(tableUnderTest.Fields())).To(BeZero())
}

func TestDbfTable_AddBooleanField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	expectedFieldName := "testBool"
	additionError := tableUnderTest.AddBooleanField(expectedFieldName)
	g.Expect(additionError).To(BeNil())

	g.Expect(tableUnderTest.NumberOfRecords()).To(BeZero())
	g.Expect(len(tableUnderTest.Fields())).To(BeNumerically("==", 1))

	addedField := tableUnderTest.Fields()[0]
	g.Expect(addedField.name).To(Equal(expectedFieldName))
	g.Expect(addedField.fieldType).To(Equal(Logical))
}

func TestDbfTable_AddBooleanField_TooLongGetsTruncated(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	expectedFieldName := "FieldName!"
	suppliedFieldName := expectedFieldName + "shouldBeTruncated"

	tableUnderTest.AddBooleanField(suppliedFieldName)

	addedField := tableUnderTest.Fields()[0]
	g.Expect(addedField.name).To(Equal(expectedFieldName))
}

func TestDbfTable_AddBooleanField_SecondAttemptFails(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	expectedFieldName := "FieldName!"

	additionError := tableUnderTest.AddBooleanField(expectedFieldName)
	g.Expect(additionError).To(BeNil())

	secondAdditionError := tableUnderTest.AddBooleanField(expectedFieldName)
	g.Expect(secondAdditionError).ToNot(BeNil())

	t.Log(secondAdditionError)
}

func TestDbfTable_AddBooleanField_ErrorAfterDataEntryStart(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	expectedFieldName := "goodField"

	additionError := tableUnderTest.AddBooleanField(expectedFieldName)
	g.Expect(additionError).To(BeNil())

	tableUnderTest.AddNewRecord()

	postDataEntryField := "badField"

	secondAdditionError := tableUnderTest.AddBooleanField(postDataEntryField)
	g.Expect(secondAdditionError).ToNot(BeNil())

	t.Log(secondAdditionError)
}

func TestDbfTable_AddDateField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	expectedFieldName := "testDate"
	additionError := tableUnderTest.AddDateField(expectedFieldName)
	g.Expect(additionError).To(BeNil())

	g.Expect(tableUnderTest.NumberOfRecords()).To(BeZero())
	g.Expect(len(tableUnderTest.Fields())).To(BeNumerically("==", 1))

	addedField := tableUnderTest.Fields()[0]
	g.Expect(addedField.name).To(Equal(expectedFieldName))
	g.Expect(addedField.fieldType).To(Equal(Date))
}

func TestDbfTable_AddTextField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	expectedFieldName := "testText"
	expectedFieldLength := byte(20)
	additionError := tableUnderTest.AddTextField(expectedFieldName, expectedFieldLength)
	g.Expect(additionError).To(BeNil())

	g.Expect(tableUnderTest.NumberOfRecords()).To(BeZero())
	g.Expect(len(tableUnderTest.Fields())).To(BeNumerically("==", 1))

	addedField := tableUnderTest.Fields()[0]
	g.Expect(addedField.name).To(Equal(expectedFieldName))
	g.Expect(addedField.fieldType).To(Equal(Character))
	g.Expect(addedField.length).To(Equal(expectedFieldLength))
}

func TestDbfTable_AddNumericField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	expectedFieldName := "testNumber"
	expectedFieldLength := byte(20)
	expectedFDecimalPlaces := byte(2)
	additionError := tableUnderTest.AddNumberField(expectedFieldName, expectedFieldLength, expectedFDecimalPlaces)
	g.Expect(additionError).To(BeNil())

	g.Expect(tableUnderTest.NumberOfRecords()).To(BeZero())
	g.Expect(len(tableUnderTest.Fields())).To(BeNumerically("==", 1))

	addedField := tableUnderTest.Fields()[0]
	g.Expect(addedField.name).To(Equal(expectedFieldName))
	g.Expect(addedField.fieldType).To(Equal(Numeric))
	g.Expect(addedField.length).To(Equal(expectedFieldLength))
	g.Expect(addedField.decimalPlaces).To(Equal(expectedFDecimalPlaces))
}

func TestDbfTable_AddFloatField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	expectedFieldName := "testFloat"
	expectedFieldLength := byte(20)
	expectedFDecimalPlaces := byte(2)
	additionError := tableUnderTest.AddFloatField(expectedFieldName, expectedFieldLength, expectedFDecimalPlaces)
	g.Expect(additionError).To(BeNil())

	g.Expect(tableUnderTest.NumberOfRecords()).To(BeZero())
	g.Expect(len(tableUnderTest.Fields())).To(BeNumerically("==", 1))

	addedField := tableUnderTest.Fields()[0]
	g.Expect(addedField.name).To(Equal(expectedFieldName))
	g.Expect(addedField.fieldType).To(Equal(Float))
	g.Expect(addedField.length).To(Equal(expectedFieldLength))
	g.Expect(addedField.decimalPlaces).To(Equal(expectedFDecimalPlaces))
}

func TestDbfTable_FieldNames(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	expectedFieldNames := []string{"first", "second"}

	for _, name := range expectedFieldNames {
		additionError := tableUnderTest.AddBooleanField(name)
		g.Expect(additionError).To(BeNil())
	}

	fieldNamesUnderTest := tableUnderTest.FieldNames()
	g.Expect(fieldNamesUnderTest).To(Equal(expectedFieldNames))
}

func TestDbfTable_DecimalPlacesInField_ValidField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	numberFieldName := "numField"
	expectedNumberDecimalPlaces := uint8(0)
	tableUnderTest.AddNumberField(numberFieldName, 5, expectedNumberDecimalPlaces)
	actualNumberDecimalPlaces, numberError := tableUnderTest.DecimalPlacesInField(numberFieldName)

	g.Expect(numberError).To(BeNil())
	g.Expect(actualNumberDecimalPlaces).To(BeNumerically("==", expectedNumberDecimalPlaces))

	floatFieldName := "floatField"
	expectedFloatDecimalPlaces := uint8(2)
	tableUnderTest.AddFloatField(floatFieldName, 10, expectedFloatDecimalPlaces)
	actualFloatDecimalPlaces, floatError := tableUnderTest.DecimalPlacesInField(floatFieldName)

	g.Expect(floatError).To(BeNil())
	g.Expect(actualFloatDecimalPlaces).To(BeNumerically("==", expectedFloatDecimalPlaces))
}

func TestDbfTable_DecimalPlacesInField_NonExistentField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	_, numberError := tableUnderTest.DecimalPlacesInField("missingField")

	g.Expect(numberError).ToNot(BeNil())
	t.Log(numberError)
}

func TestDbfTable_DecimalPlacesInField_InvalidField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	textFieldName := "textField"
	tableUnderTest.AddTextField(textFieldName, 5)
	_, numberError := tableUnderTest.DecimalPlacesInField(textFieldName)

	g.Expect(numberError).ToNot(BeNil())
	t.Log(numberError)
}

func TestDbfTable_GetRowAsSlice_InitiallyEmptyStrings(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	booldFieldName := "boolField"
	tableUnderTest.AddBooleanField(booldFieldName)

	textFieldName := "textField"
	tableUnderTest.AddBooleanField(textFieldName)

	dateFieldName := "dateField"
	tableUnderTest.AddBooleanField(dateFieldName)

	numFieldName := "numField"
	tableUnderTest.AddBooleanField(numFieldName)

	floatFieldName := "floatField"
	tableUnderTest.AddBooleanField(floatFieldName)

	recordIndex, _ := tableUnderTest.AddNewRecord()

	fieldValues := tableUnderTest.GetRowAsSlice(recordIndex)

	for _, value := range fieldValues {
		g.Expect(value).To(Equal(""))
	}
}

func TestDbfTable_GetRowAsSlice(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	boolFieldName := "boolField"
	expectedBoolFieldValue := "T"
	tableUnderTest.AddBooleanField(boolFieldName)

	textFieldName := "textField"
	expectedTextFieldValue := "some text"
	tableUnderTest.AddTextField(textFieldName, 10)

	dateFieldName := "dateField"
	expectedDateFieldValue := "20181201"
	tableUnderTest.AddDateField(dateFieldName)

	numFieldName := "numField"
	expectedNumFieldValue := "640"
	tableUnderTest.AddNumberField(numFieldName, 3, 0)

	floatFieldName := "floatField"
	expectedFloatFieldValue := "640.42"
	tableUnderTest.AddFloatField(floatFieldName, 6, 2)

	recordIndex, _ := tableUnderTest.AddNewRecord()

	tableUnderTest.SetFieldValueByName(recordIndex, boolFieldName, expectedBoolFieldValue)
	tableUnderTest.SetFieldValueByName(recordIndex, textFieldName, expectedTextFieldValue)
	tableUnderTest.SetFieldValueByName(recordIndex, dateFieldName, expectedDateFieldValue)
	tableUnderTest.SetFieldValueByName(recordIndex, numFieldName, expectedNumFieldValue)
	tableUnderTest.SetFieldValueByName(recordIndex, floatFieldName, expectedFloatFieldValue)

	fieldValues := tableUnderTest.GetRowAsSlice(recordIndex)

	g.Expect(fieldValues[0]).To(Equal(expectedBoolFieldValue))
	g.Expect(fieldValues[1]).To(Equal(expectedTextFieldValue))
	g.Expect(fieldValues[2]).To(Equal(expectedDateFieldValue))
	g.Expect(fieldValues[3]).To(Equal(expectedNumFieldValue))
	g.Expect(fieldValues[4]).To(Equal(expectedFloatFieldValue))
}

func TestDbfTable_FieldValueByName(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	boolFieldName := "boolField"
	expectedBoolFieldValue := "T"
	tableUnderTest.AddBooleanField(boolFieldName)

	textFieldName := "textField"
	expectedTextFieldValue := "some text"
	tableUnderTest.AddTextField(textFieldName, 10)

	dateFieldName := "dateField"
	expectedDateFieldValue := "20181201"
	tableUnderTest.AddDateField(dateFieldName)

	numFieldName := "numField"
	expectedNumFieldValue := "640"
	tableUnderTest.AddNumberField(numFieldName, 3, 0)

	floatFieldName := "floatField"
	expectedFloatFieldValue := "640.42"
	tableUnderTest.AddFloatField(floatFieldName, 6, 2)

	recordIndex, _ := tableUnderTest.AddNewRecord()

	tableUnderTest.SetFieldValueByName(recordIndex, boolFieldName, expectedBoolFieldValue)
	tableUnderTest.SetFieldValueByName(recordIndex, textFieldName, expectedTextFieldValue)
	tableUnderTest.SetFieldValueByName(recordIndex, dateFieldName, expectedDateFieldValue)
	tableUnderTest.SetFieldValueByName(recordIndex, numFieldName, expectedNumFieldValue)
	tableUnderTest.SetFieldValueByName(recordIndex, floatFieldName, expectedFloatFieldValue)

	g.Expect(tableUnderTest.FieldValueByName(recordIndex, boolFieldName)).To(Equal(expectedBoolFieldValue))
	g.Expect(tableUnderTest.FieldValueByName(recordIndex, textFieldName)).To(Equal(expectedTextFieldValue))
	g.Expect(tableUnderTest.FieldValueByName(recordIndex, dateFieldName)).To(Equal(expectedDateFieldValue))
	g.Expect(tableUnderTest.FieldValueByName(recordIndex, numFieldName)).To(Equal(expectedNumFieldValue))
	g.Expect(tableUnderTest.FieldValueByName(recordIndex, floatFieldName)).To(Equal(expectedFloatFieldValue))
}

func TestDbfTable_FieldValueByName_NonExistentField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	textFieldName := "textField"
	tableUnderTest.AddTextField(textFieldName, 10)

	_, valueError := tableUnderTest.FieldValueByName(0, "missingField")

	g.Expect(valueError).ToNot(BeNil())
	t.Log(valueError)
}

func TestDbfTable_SetFieldValueByName_NonExistentField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	setError := tableUnderTest.SetFieldValueByName(0, "missingField", "someText")

	g.Expect(setError).ToNot(BeNil())
	t.Log(setError)
}

func TestDbfTable_AddRecordWithNoFieldsDefined_Errors(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	recordIndex, addErr := tableUnderTest.AddNewRecord()
	g.Expect(addErr).ToNot(BeNil())
	g.Expect(recordIndex).To(BeEquivalentTo(-1))
}

func TestDbfTable_Int64FieldValueByName(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	intFieldName := "intField"
	expectedIntValue := 640
	expectedIntFieldValue := fmt.Sprintf("%d", expectedIntValue)
	tableUnderTest.AddNumberField(intFieldName, 6, 2)

	recordIndex, addErr := tableUnderTest.AddNewRecord()
	g.Expect(addErr).To(BeNil())

	tableUnderTest.SetFieldValueByName(recordIndex, intFieldName, expectedIntFieldValue)

	actualIntFieldValue, valueError := tableUnderTest.Int64FieldValueByName(recordIndex, intFieldName)

	g.Expect(valueError).To(BeNil())
	g.Expect(actualIntFieldValue).To(BeNumerically("==", expectedIntValue))
}

func TestDbfTable_Float64FieldValueByName(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	floatFieldName := "floatField"
	expectedFloatValue := 640.42
	expectedFloatFieldValue := fmt.Sprintf("%.2f", expectedFloatValue)
	tableUnderTest.AddFloatField(floatFieldName, 10, 2)

	recordIndex, addErr := tableUnderTest.AddNewRecord()
	g.Expect(addErr).To(BeNil())

	tableUnderTest.SetFieldValueByName(recordIndex, floatFieldName, expectedFloatFieldValue)

	actualFloatFieldValue, valueError := tableUnderTest.Float64FieldValueByName(recordIndex, floatFieldName)

	g.Expect(valueError).To(BeNil())
	g.Expect(actualFloatFieldValue).To(BeNumerically("==", expectedFloatValue))
}

func TestDbfTable_FieldDescriptor(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	const fieldName = "floatField"
	const fieldLength = uint8(10)
	const decimalPlaces = uint8(2)

	floatFieldName := fieldName
	tableUnderTest.AddFloatField(floatFieldName, fieldLength, decimalPlaces)

	fieldUnderTest := tableUnderTest.Fields()[0]

	g.Expect(fieldUnderTest.Name()).To(Equal(fieldName))
	g.Expect(fieldUnderTest.FieldType()).To(Equal(Float))
	g.Expect(fieldUnderTest.FieldType()).To(Equal(Float))
	g.Expect(fieldUnderTest.Length()).To(Equal(fieldLength))
	g.Expect(fieldUnderTest.DecimalPlaces()).To(Equal(decimalPlaces))

}
