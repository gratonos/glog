package iface

type Format uint8

const (
	Binary Format = iota
	Text

	formatBound
)

func (self Format) Legal() bool {
	return self < formatBound
}
