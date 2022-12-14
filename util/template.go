package util

import (
	"errors"
	"fmt"
	"strings"
)

const (
	TemplateBeginDelimiter = "{"
	TemplateEndDelimiter   = "}"
)

func ExpandTemplate(t string, lookup VariableLookup) (string, error) {
	if t == "" {
		return "", nil
	}
	if lookup == nil {
		return t, errors.New("invalid argument : VariableLookup() is nil")
	}
	var buf strings.Builder
	tokens := strings.Split(t, TemplateBeginDelimiter)
	if len(tokens) == 1 {
		return t, nil
	}
	for _, s := range tokens {
		sub := strings.Split(s, TemplateEndDelimiter)
		if len(sub) > 2 {
			return "", errors.New(fmt.Sprintf("invalid argument : token has multiple end delimiters: %v", s))
		}
		// Check case where no end delimiter is found
		if len(sub) == 1 && sub[0] == s {
			buf.WriteString(s)
			continue
		}
		// Have a valid end delimiter, so lookup the variable
		t, err := lookup(sub[0])
		if err != nil {
			return "", err
		}
		buf.WriteString(t)
		if len(sub) == 2 {
			buf.WriteString(sub[1])
		}
	}
	return buf.String(), nil
}
