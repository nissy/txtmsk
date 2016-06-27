package mask

import (
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
		return errors.New("password len 32 is over")
	case (l == 0):
		return errors.New("password is nil")
	}

	for i := l; i < n; i++ {
		pw += "*"
	}

	m.src = []byte(pw)

	return nil
}

func (m *Mask) Encrypt(text string) (string, error) {
	src := []byte(text)
	block, err := aes.NewCipher(m.src)

	if err != nil {
		return "", err
	}

	encrypted := make([]byte, aes.BlockSize+len(src))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	encryptStream := cipher.NewCTR(block, iv)
	encryptStream.XORKeyStream(encrypted[aes.BlockSize:], src)

	return base64.StdEncoding.EncodeToString(encrypted), nil
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

	decrypted := make([]byte, len(src[aes.BlockSize:]))
	decryptStream := cipher.NewCTR(block, src[:aes.BlockSize])
	decryptStream.XORKeyStream(decrypted, src[aes.BlockSize:])

	decryptedText := string(decrypted)

	if utf8.ValidString(decryptedText) {
		return decryptedText, nil
	}

	return "", errors.New("not decrypt text")
}
