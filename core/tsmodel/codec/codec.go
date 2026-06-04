package codec

import (
	"encoding/binary"
	"io"

	"github.com/hanami/tidets/commons/errors"
)

func WriteString(w io.Writer, s string) error {
	if len(s) > 0xffff {
		return commons.ErrCodecStringTooLong
	}
	if err := binary.Write(w, binary.LittleEndian, uint16(len(s))); err != nil {
		return err
	}
	_, err := w.Write([]byte(s))
	return err
}

func ReadString(r io.Reader) (string, error) {
	var n uint16
	if err := binary.Read(r, binary.LittleEndian, &n); err != nil {
		return "", err
	}
	buf := make([]byte, n)
	if _, err := io.ReadFull(r, buf); err != nil {
		return "", err
	}
	return string(buf), nil
}
