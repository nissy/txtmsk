package mask

import (
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"unicode/utf8"
)

type Mask struct {
	Password string
	key      []byte
}

func New(password string) (*Mask, error) {
	m := &Mask{
		Password: password,
	}

	if err := m.setKey(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Mask) setKey() error {
	pw := m.Password

	l := len(pw)
	n := 0

	switch {
	case (l < 16):
		n = 16
	case (l < 24):
		n = 24
	case (l < 32):
		n = 32
	case (l > 32):
		return errors.New("Password len 32 is over")
	case (l == 0):
		return errors.New("Password is nil")
	}

	for i := l; i < n; i++ {
		pw += "*"
	}

	m.key = []byte(pw)

	return nil
}

func (m *Mask) Mask(text string) (string, error) {
	if !utf8.ValidString(text) {
		return "", errors.New("Not a text")
	}

	src := compress([]byte(text))
	src, err := m.encrypt(src)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(src), nil
}

func (m *Mask) UnMask(text string) (string, error) {
	src, err := base64.StdEncoding.DecodeString(text)

	if err != nil {
		return "", err
	}

	src, err = m.decrypt(src)

	if err != nil {
		return "", err
	}

	src, err = unCompress(src)

	if err != nil {
		return "", err
	}

	dText := string(src)

	if utf8.ValidString(dText) {
		return dText, nil
	}

	return "", errors.New("Not decrypt the text")
}

func (m *Mask) encrypt(src []byte) ([]byte, error) {
	block, err := aes.NewCipher(m.key)

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
	block, err := aes.NewCipher(m.key)

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
