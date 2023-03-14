package utils

import (
	"regexp"
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
