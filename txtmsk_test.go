package txtmsk

import (
	"testing"

	"github.com/ngc224/txtmsk/mask"
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

	mText, err := tryMask(m, text)

	if err != nil {
		t.Error(err)
	}

	umText, err := tryUnMask(m, mText)

	if umText != text {
		t.Error(err)
	}
}

func TestInlineMaskToUnMask(t *testing.T) {
	m, err := mask.New(pass)

	if err != nil {
		t.Error(err)
	}

	mInlineText, err := tryMask(m, inlineText)

	if err != nil {
		t.Error(err)
	}

	umInlineText, err := tryUnMask(m, mInlineText)

	if umInlineText != inlineText {
		t.Error(err)
	}
}
