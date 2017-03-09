package txtmsk

import (
	"regexp"
	"strings"

	"github.com/nissy/txtmsk/mask"
)

var (
	inlineMarkStart = "<msk>"
	inlineMarkEnd   = "</msk>"
)

func TryMask(m *mask.Mask, text string) (string, error) {
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

func TryUnMask(m *mask.Mask, text string) (string, error) {
	umText, err := inLineUnMask(m, text)

	if err != nil {
		return "", err
	}

	if umText != text {
		return umText, nil
	}

	if !newInLineRegexp().MatchString(umText) {
		umText, err = m.UnMask(text)

		if err != nil {
			return "", err
		}
	}

	return umText, nil
}

func inLineMask(m *mask.Mask, text string) (string, error) {
	for _, v := range newInLineRegexp().FindAllStringSubmatch(text, -1) {
		line, err := m.Mask(v[1])

		if err != nil {
			return "", err
		}

		text = replaceInLine(text, v[0], line)
	}

	return text, nil
}

func inLineUnMask(m *mask.Mask, text string) (string, error) {
	for _, v := range newInLineRegexp().FindAllStringSubmatch(text, -1) {
		line, err := m.UnMask(v[1])

		if err != nil {
			continue
		}

		text = replaceInLine(text, v[0], line)
	}

	return text, nil
}

func newInLineRegexp() *regexp.Regexp {
	return regexp.MustCompile(newInlineMark(`([\s\S]+?)`))
}

func newInlineMark(element string) string {
	return inlineMarkStart + element + inlineMarkEnd
}

func replaceInLine(src, old, new string) string {
	return strings.Replace(src, old, newInlineMark(new), 1)
}

func TrimInLineTag(src string) string {
	return strings.Replace(
		strings.Replace(src, inlineMarkStart, "", -1),
		inlineMarkEnd, "", -1)
}
