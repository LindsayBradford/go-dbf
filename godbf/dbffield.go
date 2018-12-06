package godbf

// FieldDescriptor describes one field/column in a DbfTable (https://www.dbase.com/Knowledgebase/INT/db7_file_fmt.htm, Heading 1.2)
type FieldDescriptor struct {
	name          string
	fieldType     DbaseDataType
	length        byte
	decimalPlaces byte
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

// FieldType returns the length of data stored for the field
func (fd *FieldDescriptor) Length() byte {
	return fd.length
}

// FieldType returns the count of decimal places for the field
func (fd *FieldDescriptor) DecimalCount() byte {
	return fd.length
}

func (fd FieldDescriptor) usesDecimalPlaces() bool {
	return fd.fieldType.usesDecimalPlaces()
}

// A dBase Data Type (https://www.dbase.com/Knowledgebase/INT/db7_file_fmt.htm, under heading "Storage of dBASE Data Types")
type DbaseDataType byte

const (
	Character DbaseDataType = 'C'
	Logical   DbaseDataType = 'L'
	Date      DbaseDataType = 'D'
	Numeric   DbaseDataType = 'N'
	Float     DbaseDataType = 'F'
)

func (ddt DbaseDataType) Byte() byte {
	return byte(ddt)
}

func (ddt DbaseDataType) fieldLength() byte {
	switch ddt {
	case Logical:
		return 1
	case Date:
		return 8
	default:
		return 0
	}
}

func (ddt DbaseDataType) usesDecimalPlaces() bool {
	switch ddt {
	case Float, Numeric:
		return true
	default:
		return false
	}
}

func (ddt DbaseDataType) decimalPlaces() byte {
	return 0
}
