package binary

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

func appendBool(dst []byte, b bool) []byte {
	if b {
		return appendUint8(dst, 1)
	} else {
		return appendUint8(dst, 0)
	}
}

func appendUint8(dst []byte, u uint8) []byte {
	return append(dst, u)
}

func appendUint16(dst []byte, u uint16) []byte {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, u)
	return append(dst, buf...)
}

func appendUint32(dst []byte, u uint32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, u)
	return append(dst, buf...)
}

func appendUint64(dst []byte, u uint64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, u)
	return append(dst, buf...)
}

func appendFloat32(dst []byte, f float32) []byte {
	return appendUint32(dst, math.Float32bits(f))
}

func appendFloat64(dst []byte, f float64) []byte {
	return appendUint64(dst, math.Float64bits(f))
}

func appendComplex64(dst []byte, c complex64) []byte {
	dst = appendFloat32(dst, real(c))
	dst = appendFloat32(dst, imag(c))
	return dst
}

func appendComplex128(dst []byte, c complex128) []byte {
	dst = appendFloat64(dst, real(c))
	dst = appendFloat64(dst, imag(c))
	return dst
}

func appendString(dst []byte, str string) []byte {
	size := len(str)
	if size > math.MaxUint16 {
		size = math.MaxUint16
	}
	dst = appendUint16(dst, uint16(size))
	dst = append(dst, str[:size]...)
	return dst
}

func appendShortString(dst []byte, str string) []byte {
	size := len(str)
	if size > math.MaxUint8 {
		size = math.MaxUint8
	}
	dst = appendUint8(dst, uint8(size))
	dst = append(dst, str[:size]...)
	return dst
}

func readBool(reader *bufio.Reader) (bool, error) {
	u, err := readUint8(reader)
	if err != nil {
		return false, err
	}
	if u > 1 {
		return false, newFormatError(fmt.Sprintf("illegal bool value: %d", u))
	}
	if u == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func readUint8(reader *bufio.Reader) (uint8, error) {
	buf := make([]byte, 1)
	if err := read(buf, reader); err != nil {
		return 0, err
	} else {
		return buf[0], nil
	}
}

func readUint16(reader *bufio.Reader) (uint16, error) {
	buf := make([]byte, 2)
	if err := read(buf, reader); err != nil {
		return 0, err
	} else {
		return binary.LittleEndian.Uint16(buf), nil
	}
}

func readUint32(reader *bufio.Reader) (uint32, error) {
	buf := make([]byte, 4)
	if err := read(buf, reader); err != nil {
		return 0, err
	} else {
		return binary.LittleEndian.Uint32(buf), nil
	}
}

func readUint64(reader *bufio.Reader) (uint64, error) {
	buf := make([]byte, 8)
	if err := read(buf, reader); err != nil {
		return 0, err
	} else {
		return binary.LittleEndian.Uint64(buf), nil
	}
}

func readFloat32(reader *bufio.Reader) (float32, error) {
	u, err := readUint32(reader)
	if err != nil {
		return 0.0, err
	}
	return math.Float32frombits(u), nil
}

func readFloat64(reader *bufio.Reader) (float64, error) {
	u, err := readUint64(reader)
	if err != nil {
		return 0.0, err
	}
	return math.Float64frombits(u), nil
}

func readComplex64(reader *bufio.Reader) (complex64, error) {
	r, err := readFloat32(reader)
	if err != nil {
		return 0.0, err
	}
	i, err := readFloat32(reader)
	if err != nil {
		return 0.0, err
	}
	return complex(r, i), nil
}

func readComplex128(reader *bufio.Reader) (complex128, error) {
	r, err := readFloat64(reader)
	if err != nil {
		return 0.0, err
	}
	i, err := readFloat64(reader)
	if err != nil {
		return 0.0, err
	}
	return complex(r, i), nil
}

func readString(reader *bufio.Reader) (string, error) {
	size, err := readUint16(reader)
	if err != nil {
		return "", err
	}
	return readStr(reader, uint(size))
}

func readShortString(reader *bufio.Reader) (string, error) {
	size, err := readUint8(reader)
	if err != nil {
		return "", err
	}
	return readStr(reader, uint(size))
}

func readStr(reader *bufio.Reader, size uint) (string, error) {
	buf := make([]byte, size)
	if err := read(buf, reader); err != nil {
		return "", err
	}
	return string(buf), nil
}

func read(buf []byte, reader *bufio.Reader) error {
	_, err := io.ReadFull(reader, buf)
	if err != nil {
		return newIOError(err)
	}
	return nil
}
