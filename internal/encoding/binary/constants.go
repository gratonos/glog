package binary

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
	endIdentity
)

type valueKind uint8

const (
	boolValue valueKind = iota
	byteValue
	runeValue
	intValue
	uintValue
	uintptrValue
	floatValue
	complexValue
	stringValue
	timeValue
	durationValue
)

const binaryVersion = 0

const (
	sizeOfVersion = 1
	sizeOfLoadLen = 4
	sizeOfHeader  = sizeOfVersion + sizeOfLoadLen
)
