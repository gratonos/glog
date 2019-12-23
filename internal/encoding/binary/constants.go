package binary

import (
	"math"
)

type FieldKind uint8

const (
	Timestamp FieldKind = iota
	Level
	Pkg
	Func
	File
	Line
	Mark
	Msg
	KeyValue

	MaxFieldKind = KeyValue

	End FieldKind = math.MaxUint8
)

type ValueKind uint8

const (
	Bool ValueKind = iota
	Byte
	Rune
	Int8
	Int16
	Int32
	Int64
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	String
	Time
	Duration

	MaxValueKind = Duration

	IllegalValueKind ValueKind = math.MaxUint8
)

var binaryMagic = []byte{0x14, 0xf2, 0x79, 0xd3, 0x6b, 0xe7, 0x3d}

const binaryVersion = 0

const (
	sizeOfMagic   = 7
	sizeOfVersion = 1
	sizeOfHeader  = sizeOfMagic + sizeOfVersion
)
