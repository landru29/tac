package main

import (
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/landru29/tac/internal/base45"
	"github.com/liyue201/goqr"
)

func decodeQR(imgStream io.Reader) (io.Reader, error) {
	img, _, err := image.Decode(imgStream)
	if err != nil {
		return nil, err
	}

	qrCodes, err := goqr.Recognize(img)
	if err != nil {
		return nil, err
	}

	if len(qrCodes) != 1 {
		return nil, errors.New("should find one QR code on the image")
	}

	data := []byte(qrCodes[0].Payload)[4:]

	return bytes.NewReader(data), nil
}

func main() {
	f, err := os.Open("testdata/tac.jpg")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = f.Close()
	}()

	b45, err := decodeQR(f)
	if err != nil {
		log.Fatal(err)
	}

	bin, err := base45.Decode(b45)
	if err != nil {
		log.Fatal(err)
	}

	uncompressed, err := zlib.NewReader(bin)
	if err != nil {
		log.Fatal(err)
	}

	out, err := ioutil.ReadAll(uncompressed)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}
