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
	src      []byte
}

func New(password string) (*Mask, error) {
	m := &Mask{
		Password: password,
	}

	if err := m.setSrc(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Mask) setSrc() error {
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

	m.src = []byte(pw)

	return nil
}

func (m *Mask) Encrypt(text string) (string, error) {
	if !utf8.ValidString(text) {
		return "", errors.New("Not a text")
	}

	src := compress([]byte(text))

	block, err := aes.NewCipher(m.src)

	if err != nil {
		return "", err
	}

	encSrc := make([]byte, aes.BlockSize+len(src))
	iv := encSrc[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	encStream := cipher.NewCTR(block, iv)
	encStream.XORKeyStream(encSrc[aes.BlockSize:], src)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encSrc), nil
}

func (m *Mask) Decrypt(text string) (string, error) {
	src, err := base64.StdEncoding.DecodeString(text)

	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(m.src)

	if err != nil {
		return "", err
	}

	decSrc := make([]byte, len(src[aes.BlockSize:]))
	decStream := cipher.NewCTR(block, src[:aes.BlockSize])
	decStream.XORKeyStream(decSrc, src[aes.BlockSize:])

	decUnCompSrc, err := unCompress(decSrc)

	if err != nil {
		return "", err
	}

	decText := string(decUnCompSrc)

	if utf8.ValidString(decText) {
		return decText, nil
	}

	return "", errors.New("Not decrypt the text")
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
