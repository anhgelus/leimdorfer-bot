package utils

import (
	"regexp"
	"strings"
)

type CaseInsensitiveReplacer struct {
	toReplace   *regexp.Regexp
	replaceWith string
}

func NewCaseInsensitiveReplacer(toReplace, replaceWith string) *CaseInsensitiveReplacer {
	return &CaseInsensitiveReplacer{
		toReplace:   regexp.MustCompile("(?i)" + toReplace),
		replaceWith: replaceWith,
	}
}

func (cir *CaseInsensitiveReplacer) Replace(str string) string {
	return cir.toReplace.ReplaceAllString(str, cir.replaceWith)
}

func Replace(cir *CaseInsensitiveReplacer, str string) string {
	return cir.toReplace.ReplaceAllString(str, cir.replaceWith)
}

func SpeedRemover(msg string, toRemove string) string {
	final := msg
	if strings.Contains(final, toRemove) {
		final = strings.ReplaceAll(final, toRemove, "")
	}
	return final
}

func SpeedRemoverRegex(msg string, re *regexp.Regexp) string {
	return re.ReplaceAllString(msg, "")
}
