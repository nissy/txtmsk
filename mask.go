package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ngc224/txtmsk/mask"
)

func runMask(m *mask.Mask, text string, tagname string) error {
	mText, err := inLineMask(m, text, tagname)

	if err != nil {
		return err
	}

	if mText != text {
		fmt.Println(mText)
		return nil
	}

	mText, err = m.Mask(text)

	if err != nil {
		return err
	}

	fmt.Println(mText)
	return nil
}

func runUnMask(m *mask.Mask, text string, tagname string) error {
	umText, err := inLineUnMask(m, text, tagname)

	if err != nil {
		return err
	}

	if umText != text {
		fmt.Println(umText)
		return nil
	}

	umText, err = m.UnMask(text)

	if err != nil {
		return err
	}

	fmt.Println(umText)
	return nil
}

func inLineMask(m *mask.Mask, text string, tagname string) (string, error) {
	for _, v := range newInLineRegexp(tagname).FindAllStringSubmatch(text, -1) {
		mLine, err := m.Mask(v[1])

		if err != nil {
			return "", err
		}

		text = replaceInLine(text, v[0], mLine, tagname)
	}

	return text, nil
}

func inLineUnMask(m *mask.Mask, text string, tagname string) (string, error) {
	for _, v := range newInLineRegexp(tagname).FindAllStringSubmatch(text, -1) {
		umLine, err := m.UnMask(v[1])

		if err != nil {
			return "", err
		}

		text = replaceInLine(text, v[0], umLine, tagname)
	}

	return text, nil
}

func newInLineRegexp(tagname string) *regexp.Regexp {
	return regexp.MustCompile(`<` + tagname + `>([\s\S]+?)</` + tagname + `>`)
}

func replaceInLine(src, old, new, tagname string) string {
	return strings.Replace(src, old, "<"+tagname+">"+new+"</"+tagname+">", 1)
}
