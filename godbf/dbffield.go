package godbf

type DbfField struct {
	fieldName          string
	fieldType          dbfFieldType
	fieldLength        byte
	fieldDecimalPlaces byte
	fieldStore         [32]byte
}

func (df DbfField) usesDecimalPlaces() bool {
	return df.fieldType.usesDecimalPlaces()
}

type dbfFieldType byte

const (
	Character dbfFieldType = 'C'
	Logical   dbfFieldType = 'L'
	Date      dbfFieldType = 'D'
	Numeric   dbfFieldType = 'N'
	Float     dbfFieldType = 'F'
)

func (ddt dbfFieldType) Byte() byte {
	return byte(ddt)
}

func (ddt dbfFieldType) fieldLength() byte {
	switch ddt {
	case Logical:
		return 1
	case Date:
		return 8
	default:
		return 0
	}
}

func (ddt dbfFieldType) usesDecimalPlaces() bool {
	switch ddt {
	case Float, Numeric:
		return true
	default:
		return false
	}
}

func (ddt dbfFieldType) decimalPlaces() byte {
	return 0
}
