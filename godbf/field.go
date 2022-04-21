package godbf

// FieldDescriptor describes one field/column in a DbfTable as per https://www.dbase.com/Knowledgebase/INT/db7_file_fmt.htm, Heading 1.2.
type FieldDescriptor struct {
	name          string
	fieldType     DbaseDataType
	length        byte
	decimalPlaces byte // Field decimal count in binary
	fieldStore    [32]byte
}

// Name returns the column name of the field
func (fd *FieldDescriptor) Name() string {
	return fd.name
}

// FieldType returns the type of data stored for the field as a DbaseDataType
func (fd *FieldDescriptor) FieldType() DbaseDataType {
	return fd.fieldType
}

// Length returns the length of data stored for the field
func (fd *FieldDescriptor) Length() byte {
	return fd.length
}

// DecimalPlaces returns the count of decimal places for the field
func (fd *FieldDescriptor) DecimalPlaces() byte {
	return fd.decimalPlaces
}

func (fd FieldDescriptor) usesDecimalPlaces() bool {
	return fd.fieldType.usesDecimalCount()
}

// DbaseDataType is dBase data type, as per https://www.dbase.com/Knowledgebase/INT/db7_file_fmt.htm, under heading "Storage of dBASE Data Types".
type DbaseDataType byte

const (
	Character DbaseDataType = 'C'
	Logical   DbaseDataType = 'L'
	Date      DbaseDataType = 'D'
	Numeric   DbaseDataType = 'N'
	Float     DbaseDataType = 'F'
)

func (ddt DbaseDataType) byte() byte {
	return byte(ddt)
}

const notApplicable = 0x00

// fixedFieldLength returns the length in bytes in for the data type if it describes a fixed-length field.
func (ddt DbaseDataType) fixedFieldLength() byte {
	switch ddt {
	case Logical:
		return 1
	case Date:
		return 8
	default:
		return notApplicable
	}
}

// usesDecimalCount indicates whether the data type describes a field that makes use of a field's decimal count setting.
func (ddt DbaseDataType) usesDecimalCount() bool {
	switch ddt {
	case Float, Numeric:
		return true
	default:
		return false
	}
}

// decimalCountNotApplicable is a convenience decorator supplying a 0-valued byte. THis is used indicate that the data
// type describes a field that does not make use its decimal count setting.
func (ddt DbaseDataType) decimalCountNotApplicable() byte {
	return notApplicable
}
