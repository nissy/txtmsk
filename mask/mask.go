package mask

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type Mask struct {
	Password string
	src      []byte
}

func NewMask(password string) (*Mask, error) {
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
		return errors.New("Error: password len 32 is over")
	case (l == 0):
		return errors.New("Error: password is nil")
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

	cipher_text := make([]byte, aes.BlockSize+len(src))
	iv := cipher_text[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	encrypt_stream := cipher.NewCTR(block, iv)
	encrypt_stream.XORKeyStream(cipher_text[aes.BlockSize:], src)

	return base64.StdEncoding.EncodeToString(cipher_text), nil
}

func (m *Mask) Decrypt(text string) (string, error) {
	cipher_text, err := base64.StdEncoding.DecodeString(text)

	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(m.src)

	if err != nil {
		return "", err
	}

	decrypted_text := make([]byte, len(cipher_text[aes.BlockSize:]))
	decrypt_stream := cipher.NewCTR(block, cipher_text[:aes.BlockSize])
	decrypt_stream.XORKeyStream(decrypted_text, cipher_text[aes.BlockSize:])

	return string(decrypted_text), nil
}
