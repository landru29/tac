package base45

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"strings"
)

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:"

func decode(rd io.Reader) ([]byte, error) {
	dataC := []byte{0, 0, 0}
	dataI := []int{0, 0, 0}

	n, err := rd.Read(dataC)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, io.EOF
	}

	for i, c := range dataC {
		if c == 0 {
			continue
		}
		dataI[i] = strings.Index(alphabet, string(c))
		if dataI[i] < 0 {
			return nil, errors.New("invalid character")
		}
	}

	if n == 1 {
		return []byte{byte(dataI[0])}, nil
	}

	if n == 2 {
		result := dataI[0] + len(alphabet)*dataI[1]
		a := make([]byte, 2)
		binary.BigEndian.PutUint16(a, uint16(result))
		return []byte{a[1]}, nil
	}

	result := dataI[0] + len(alphabet)*(dataI[1]+len(alphabet)*dataI[2])
	a := make([]byte, 2)
	binary.BigEndian.PutUint16(a, uint16(result))

	return a, nil
}

func Decode(rd io.Reader) (io.Reader, error) {
	result := []byte{}
	for {
		data, err := decode(rd)
		if err == io.EOF {
			return bytes.NewReader(result), nil
		}
		if err != nil {
			return nil, err
		}

		result = append(result, data...)
	}
}
