package mask

import (
	"errors"
	"testing"
)

var (
	text = "aaaaaaaaaaaaaaaaaaaaaa"
	pass = "test"
)

func TestMaskToUnMask(t *testing.T) {
	m, err := New(pass)

	if err != nil {
		t.Error(err)
	}

	mText, err := m.Mask(text)

	if err != nil {
		t.Error(err)
	}

	umText, err := m.UnMask(mText)

	if err != nil {
		t.Error(err)
	}

	if text != umText {
		t.Error(errors.New("unmask text no match"))
	}
}
