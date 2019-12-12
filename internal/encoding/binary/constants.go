package binary

import (
	"math"
)

type fieldKind uint8

const (
	timeField fieldKind = iota
	levelField
	pkgField
	funcField
	fileField
	lineField
	markField
	msgField
	keyValueField

	endIdentifier fieldKind = math.MaxUint8
)

type valueKind uint8

const (
	boolValue valueKind = iota
	byteValue
	runeValue
	int8Value
	int16Value
	int32Value
	int64Value
	uint8Value
	uint16Value
	uint32Value
	uint64Value
	uintptrValue // coerced into uint64
	float32Value
	float64Value
	complex64Value
	complex128Value
	stringValue
	timeValue     // int64
	durationValue // int64
)

const binaryVersion = 0

const (
	sizeOfVersion = 1
	sizeOfLoadLen = 4
	sizeOfHeader  = sizeOfVersion + sizeOfLoadLen
)
