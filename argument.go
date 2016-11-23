package osc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// Argument represents an OSC argument.
// An OSC argument can have many different types, which is why
// we choose to represent them with an interface.
type Argument interface {
	io.WriterTo

	Equal(Argument) bool
	ReadInt32() (int32, error)
	ReadFloat32() (float32, error)
	ReadBool() (bool, error)
	ReadString() (string, error)
	ReadBlob() ([]byte, error)
	Typetag() byte
}

// Int represents a 32-bit integer.
type Int int32

// Equal returns true if the argument equals the other one, false otherwise.
func (i Int) Equal(other Argument) bool {
	if other.Typetag() != TypetagInt {
		return false
	}
	i2 := other.(Int)
	return i == i2
}

// ReadInt32 reads a 32-bit integer from the arg.
func (i Int) ReadInt32() (int32, error) { return int32(i), nil }

// ReadFloat32 reads a 32-bit float from the arg.
func (i Int) ReadFloat32() (float32, error) { return 0, ErrInvalidTypeTag }

// ReadBool bool reads a boolean from the arg.
func (i Int) ReadBool() (bool, error) { return false, ErrInvalidTypeTag }

// ReadString string reads a string from the arg.
func (i Int) ReadString() (string, error) { return "", ErrInvalidTypeTag }

// ReadBlob reads a slice of bytes from the arg.
func (i Int) ReadBlob() ([]byte, error) { return nil, ErrInvalidTypeTag }

// Typetag returns the argument's type tag.
func (i Int) Typetag() byte { return TypetagInt }

// WriteTo writes the arg to an io.Writer.
func (i Int) WriteTo(w io.Writer) (int64, error) {
	written, err := fmt.Fprintf(w, "%d", i)
	return int64(written), err
}

// Float represents a 32-bit float.
type Float float32

// Equal returns true if the argument equals the other one, false otherwise.
func (f Float) Equal(other Argument) bool {
	if other.Typetag() != TypetagFloat {
		return false
	}
	f2 := other.(Float)
	return f == f2
}

// ReadInt32 reads a 32-bit integer from the arg.
func (f Float) ReadInt32() (int32, error) { return 0, ErrInvalidTypeTag }

// ReadFloat32 reads a 32-bit float from the arg.
func (f Float) ReadFloat32() (float32, error) { return float32(f), nil }

// ReadBool bool reads a boolean from the arg.
func (f Float) ReadBool() (bool, error) { return false, ErrInvalidTypeTag }

// ReadString string reads a string from the arg.
func (f Float) ReadString() (string, error) { return "", ErrInvalidTypeTag }

// ReadBlob reads a slice of bytes from the arg.
func (f Float) ReadBlob() ([]byte, error) { return nil, ErrInvalidTypeTag }

// Typetag returns the argument's type tag.
func (f Float) Typetag() byte { return TypetagFloat }

// WriteTo writes the arg to an io.Writer.
func (f Float) WriteTo(w io.Writer) (int64, error) {
	written, err := fmt.Fprintf(w, "%f", f)
	return int64(written), err
}

// Bool represents a boolean value.
type Bool bool

// Equal returns true if the argument equals the other one, false otherwise.
func (b Bool) Equal(other Argument) bool {
	if other.Typetag() != TypetagFalse && other.Typetag() != TypetagTrue {
		return false
	}
	b2 := other.(Bool)
	return b == b2
}

// ReadInt32 reads a 32-bit integer from the arg.
func (b Bool) ReadInt32() (int32, error) { return 0, ErrInvalidTypeTag }

// ReadFloat32 reads a 32-bit float from the arg.
func (b Bool) ReadFloat32() (float32, error) { return 0, ErrInvalidTypeTag }

// ReadBool bool reads a boolean from the arg.
func (b Bool) ReadBool() (bool, error) { return bool(b), nil }

// ReadString string reads a string from the arg.
func (b Bool) ReadString() (string, error) { return "", ErrInvalidTypeTag }

// ReadBlob reads a slice of bytes from the arg.
func (b Bool) ReadBlob() ([]byte, error) { return nil, ErrInvalidTypeTag }

// Typetag returns the argument's type tag.
func (b Bool) Typetag() byte {
	if bool(b) {
		return TypetagTrue
	}
	return TypetagFalse
}

// WriteTo writes the arg to an io.Writer.
func (b Bool) WriteTo(w io.Writer) (int64, error) {
	written, err := fmt.Fprintf(w, "%t", b)
	return int64(written), err
}

// String is a string.
type String string

// Equal returns true if the argument equals the other one, false otherwise.
func (s String) Equal(other Argument) bool {
	if other.Typetag() != TypetagString {
		return false
	}
	s2 := other.(String)
	return s == s2
}

// ReadInt32 reads a 32-bit integer from the arg.
func (s String) ReadInt32() (int32, error) { return 0, ErrInvalidTypeTag }

// ReadFloat32 reads a 32-bit float from the arg.
func (s String) ReadFloat32() (float32, error) { return 0, ErrInvalidTypeTag }

// ReadBool bool reads a boolean from the arg.
func (s String) ReadBool() (bool, error) { return false, ErrInvalidTypeTag }

// ReadString string reads a string from the arg.
func (s String) ReadString() (string, error) { return string(s), nil }

// ReadBlob reads a slice of bytes from the arg.
func (s String) ReadBlob() ([]byte, error) { return nil, ErrInvalidTypeTag }

// Typetag returns the argument's type tag.
func (s String) Typetag() byte { return TypetagString }

// WriteTo writes the arg to an io.Writer.
func (s String) WriteTo(w io.Writer) (int64, error) {
	written, err := fmt.Fprintf(w, "%s", s)
	return int64(written), err
}

// Blob is a slice of bytes.
type Blob []byte

func (b Blob) Equal(other Argument) bool {
	if other.Typetag() != TypetagBlob {
		return false
	}
	b2 := other.(Blob)
	if len(b) != len(b2) {
		return false
	}
	return bytes.Equal(b, b2)
}

// ReadInt32 reads a 32-bit integer from the arg.
func (b Blob) ReadInt32() (int32, error) { return 0, ErrInvalidTypeTag }

// ReadFloat32 reads a 32-bit float from the arg.
func (b Blob) ReadFloat32() (float32, error) { return 0, ErrInvalidTypeTag }

// ReadBool bool reads a boolean from the arg.
func (b Blob) ReadBool() (bool, error) { return false, ErrInvalidTypeTag }

// ReadString string reads a string from the arg.
func (b Blob) ReadString() (string, error) { return "", ErrInvalidTypeTag }

// ReadBlob reads a slice of bytes from the arg.
func (b Blob) ReadBlob() ([]byte, error) { return []byte(b), nil }

// Typetag returns the argument's type tag.
func (b Blob) Typetag() byte { return TypetagBlob }

// WriteTo writes the arg to an io.Writer.
func (b Blob) WriteTo(w io.Writer) (int64, error) {
	written, err := w.Write([]byte(b))
	return int64(written), err
}

// ParseArgument parses an OSC message argument given a type tag and some data.
func ParseArgument(tt byte, data []byte) (Argument, int64, error) {
	switch tt {
	case TypetagInt:
		var val int32
		_ = binary.Read(bytes.NewReader(data), byteOrder, &val) // Never fails
		return Int(val), 4, nil
	case TypetagFloat:
		var val float32
		_ = binary.Read(bytes.NewReader(data), byteOrder, &val) // Never fails
		return Float(val), 4, nil
	case TypetagTrue:
		return Bool(true), 0, nil
	case TypetagFalse:
		return Bool(false), 0, nil
	case TypetagString:
		s, idx := ReadString(data)
		return String(s), idx, nil
	case TypetagBlob:
		var length int32
		if err := binary.Read(bytes.NewReader(data), byteOrder, &length); err != nil {
			return nil, 0, err
		}
		b, bl := ReadBlob(length, data[4:])
		return Blob(b), bl + 4, nil
	default:
		return nil, 0, ErrInvalidTypeTag
	}
}