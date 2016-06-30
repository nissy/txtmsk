package main

import (
	"testing"

	"github.com/ngc224/txtmsk/mask"
)

var (
	text       = "aaaaaaaaaaaaaaaaaaaaaa"
	inlineText = "a" + inlineStartTag + "b" + inlineEndTag + "c" + inlineStartTag + "d" + inlineEndTag + "e"
	pass       = "test"
)

func TestMaskToUnMask(t *testing.T) {
	m, err := mask.New(pass)

	if err != nil {
		t.Error(err)
	}

	mText, err := runMask(m, text)

	if err != nil {
		t.Error(err)
	}

	umText, err := runUnMask(m, mText)

	if umText != text {
		t.Error(err)
	}
}

func TestInlineMaskToUnMask(t *testing.T) {
	m, err := mask.New(pass)

	if err != nil {
		t.Error(err)
	}

	mInlineText, err := runMask(m, inlineText)

	if err != nil {
		t.Error(err)
	}

	umInlineText, err := runUnMask(m, mInlineText)

	if umInlineText != inlineText {
		t.Error(err)
	}
}
