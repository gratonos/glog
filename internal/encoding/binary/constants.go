package binary

type fieldKind uint8

const (
	timestampField fieldKind = iota
	levelField
	pkgField
	fileNameField
	fileLineField
	markField
	msgField
	keyValuePairField
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
)

const binaryVersion = 0

const (
	sizeOfVersion = 1
	sizeOfLoadLen = 4
	sizeOfHeader  = sizeOfVersion + sizeOfLoadLen
)
