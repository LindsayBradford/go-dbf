package godbf

import (
	"testing"

	. "github.com/onsi/gomega"
)

const testEncoding = "UTF-8"

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
	g.Expect(addedField.fieldName).To(Equal(expectedFieldName))
	g.Expect(addedField.fieldType).To(Equal("L"))
}

func TestDbfTable_AddBooleanField_TooLongGetsTruncated(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	expectedFieldName := "FieldName!"
	suppliedFieldName := expectedFieldName + "shouldBeTruncated"

	tableUnderTest.AddBooleanField(suppliedFieldName)

	addedField := tableUnderTest.Fields()[0]
	g.Expect(addedField.fieldName).To(Equal(expectedFieldName))
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
	g.Expect(addedField.fieldName).To(Equal(expectedFieldName))
	g.Expect(addedField.fieldType).To(Equal("D"))
}

func TestDbfTable_AddTextField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	expectedFieldName := "testText"
	expectedFieldLength := uint8(20)
	additionError := tableUnderTest.AddTextField(expectedFieldName, expectedFieldLength)
	g.Expect(additionError).To(BeNil())

	g.Expect(tableUnderTest.NumberOfRecords()).To(BeZero())
	g.Expect(len(tableUnderTest.Fields())).To(BeNumerically("==", 1))

	addedField := tableUnderTest.Fields()[0]
	g.Expect(addedField.fieldName).To(Equal(expectedFieldName))
	g.Expect(addedField.fieldType).To(Equal("C"))
	g.Expect(addedField.fieldLength).To(Equal(expectedFieldLength))
}

func TestDbfTable_AddNumberField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	expectedFieldName := "testNumber"
	expectedFieldLength := uint8(20)
	expectedFDecimalPlaces := uint8(2)
	additionError := tableUnderTest.AddNumberField(expectedFieldName, expectedFieldLength, expectedFDecimalPlaces)
	g.Expect(additionError).To(BeNil())

	g.Expect(tableUnderTest.NumberOfRecords()).To(BeZero())
	g.Expect(len(tableUnderTest.Fields())).To(BeNumerically("==", 1))

	addedField := tableUnderTest.Fields()[0]
	g.Expect(addedField.fieldName).To(Equal(expectedFieldName))
	g.Expect(addedField.fieldType).To(Equal("N"))
	g.Expect(addedField.fieldLength).To(Equal(expectedFieldLength))
	g.Expect(addedField.fieldDecimalPlaces).To(Equal(expectedFDecimalPlaces))
}

func TestDbfTable_AddFloatField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	expectedFieldName := "testFloat"
	expectedFieldLength := uint8(20)
	expectedFDecimalPlaces := uint8(2)
	additionError := tableUnderTest.AddFloatField(expectedFieldName, expectedFieldLength, expectedFDecimalPlaces)
	g.Expect(additionError).To(BeNil())

	g.Expect(tableUnderTest.NumberOfRecords()).To(BeZero())
	g.Expect(len(tableUnderTest.Fields())).To(BeNumerically("==", 1))

	addedField := tableUnderTest.Fields()[0]
	g.Expect(addedField.fieldName).To(Equal(expectedFieldName))
	g.Expect(addedField.fieldType).To(Equal("F"))
	g.Expect(addedField.fieldLength).To(Equal(expectedFieldLength))
	g.Expect(addedField.fieldDecimalPlaces).To(Equal(expectedFDecimalPlaces))
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
