package main

import (
	"regexp"
	"strings"

	"github.com/ngc224/txtmsk/mask"
)

var (
	markName        = "msk"
	inlineMarkStart = "<" + markName + ">"
	inlineMarkEnd   = "</" + markName + ">"
)

func tryMask(m *mask.Mask, text string) (string, error) {
	mText, err := inLineMask(m, text)

	if err != nil {
		return "", err
	}

	if mText != text {
		return mText, nil
	}

	mText, err = m.Mask(text)

	if err != nil {
		return "", err
	}

	return mText, nil
}

func tryUnMask(m *mask.Mask, text string) (string, error) {
	umText, err := inLineUnMask(m, text)

	if err != nil {
		return "", err
	}

	if umText != text {
		return umText, nil
	}

	umText, err = m.UnMask(text)

	if err != nil {
		return "", err
	}

	return umText, nil
}

func inLineMask(m *mask.Mask, text string) (string, error) {
	for _, v := range newInLineRegexp().FindAllStringSubmatch(text, -1) {
		mLine, err := m.Mask(v[1])

		if err != nil {
			return "", err
		}

		text = replaceInLine(text, v[0], mLine)
	}

	return text, nil
}

func inLineUnMask(m *mask.Mask, text string) (string, error) {
	for _, v := range newInLineRegexp().FindAllStringSubmatch(text, -1) {
		umLine, err := m.UnMask(v[1])

		if err != nil {
			return "", err
		}

		text = replaceInLine(text, v[0], umLine)
	}

	return text, nil
}

func newInLineRegexp() *regexp.Regexp {
	return regexp.MustCompile(newInlineMark(`([\s\S]+?)`))
}

func replaceInLine(src, old, new string) string {
	return strings.Replace(src, old, newInlineMark(new), 1)
}

func newInlineMark(element string) string {
	return inlineMarkStart + element + inlineMarkEnd
}

func trimInLineTag(src string) string {
	return strings.Replace(
		strings.Replace(src, inlineMarkStart, "", -1),
		inlineMarkEnd, "", -1)
}
