package mask

import (
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"unicode/utf8"
)

var (
	ErrNotDecrypt = errors.New("Not decrypt the text")
	ErrNotMasked  = errors.New("Not masked the text")
	ErrNotUseText = errors.New("This is not a text of utf8")
	ErrEmptyText  = errors.New("Text is empty")
)

type Mask struct {
	password string
	realkey  []byte
}

func New(password string) (*Mask, error) {
	realkey, err := newRealkey(password)

	if err != nil {
		return nil, err
	}

	return &Mask{
		password: password,
		realkey:  realkey,
	}, nil
}

func newRealkey(key string) ([]byte, error) {
	l := len(key)
	n := 0

	switch {
	case (l < 32):
		n = 32
	case (l > 32):
		return nil, errors.New("Password len 32 is over")
	case (l == 0):
		return nil, errors.New("Password is null")
	}

	for i := l; i < n; i++ {
		key += "*"
	}

	return []byte(key), nil
}

func (m *Mask) Mask(text string) (string, error) {
	if err := verifyText(text); err != nil {
		return "", err
	}

	src := compress([]byte(text))
	src, err := m.encrypt(src)

	if err != nil {
		return "", err
	}

	return base64.RawStdEncoding.EncodeToString(src), nil
}

func (m *Mask) UnMask(text string) (string, error) {
	if err := verifyText(text); err != nil {
		return "", err
	}

	if err := verifyMaskText(text); err != nil {
		return "", err
	}

	src, err := base64.RawStdEncoding.DecodeString(text)

	if err != nil {
		return "", ErrNotMasked
	}

	//TODO panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", ErrNotDecrypt)
			os.Exit(1)
		}
	}()

	src, err = m.decrypt(src)

	if err != nil {
		return "", err
	}

	src, err = unCompress(src)

	if err != nil {
		if err == zlib.ErrHeader {
			return "", ErrNotDecrypt
		}

		return "", err
	}

	dText := string(src)

	if err := verifyText(dText); err != nil {
		return "", err
	}

	return dText, nil
}

func (m *Mask) encrypt(src []byte) ([]byte, error) {
	block, err := aes.NewCipher(m.realkey)

	if err != nil {
		return nil, err
	}

	eSrc := make([]byte, aes.BlockSize+len(src))
	iv := eSrc[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(eSrc[aes.BlockSize:], src)

	if err != nil {
		return nil, err
	}

	return eSrc, nil
}

func (m *Mask) decrypt(src []byte) ([]byte, error) {
	block, err := aes.NewCipher(m.realkey)

	if err != nil {
		return nil, err
	}

	dSrc := make([]byte, len(src[aes.BlockSize:]))
	stream := cipher.NewCTR(block, src[:aes.BlockSize])
	stream.XORKeyStream(dSrc, src[aes.BlockSize:])

	return dSrc, nil
}

func compress(src []byte) []byte {
	var buf bytes.Buffer

	w := zlib.NewWriter(&buf)
	w.Write(src)
	w.Close()

	return buf.Bytes()
}

func unCompress(src []byte) ([]byte, error) {
	var buf bytes.Buffer
	var dstBuf bytes.Buffer

	buf.Write(src)

	r, err := zlib.NewReader(&buf)

	if err != nil {
		return nil, err
	}

	io.Copy(&dstBuf, r)
	r.Close()

	return dstBuf.Bytes(), nil
}

func verifyMaskText(text string) error {
	if len(text) < 39 {
		return ErrNotMasked
	}

	if !regexp.MustCompile(`^[A-Za-z0-9/+]*=*$`).MatchString(text) {
		return ErrNotMasked
	}

	return nil
}

func verifyText(text string) error {
	if !utf8.ValidString(text) {
		return ErrNotUseText
	}

	if len(text) == 0 {
		return ErrEmptyText
	}

	return nil
}
