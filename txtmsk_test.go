package txtmsk

import (
	"testing"

	"github.com/nissy/txtmsk/mask"
)

var (
	text       = "aaaaaaaaaaaaaaaaaaaaaa"
	inlineText = "a" + newInlineMark("b") + "c" + newInlineMark("d") + "e"
	pass       = "test"
)

func TestMaskToUnMask(t *testing.T) {
	m, err := mask.New(pass)

	if err != nil {
		t.Error(err)
	}

	mText, err := TryMask(m, text)

	if err != nil {
		t.Error(err)
	}

	umText, err := TryUnMask(m, mText)

	if umText != text {
		t.Error(err)
	}
}

func TestInlineMaskToUnMask(t *testing.T) {
	m, err := mask.New(pass)

	if err != nil {
		t.Error(err)
	}

	mInlineText, err := TryMask(m, inlineText)

	if err != nil {
		t.Error(err)
	}

	umInlineText, err := TryUnMask(m, mInlineText)

	if umInlineText != inlineText || err != nil {
		t.Error(err)
	}
}
