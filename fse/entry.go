package fse

import (
	"errors"
	"strings"
)

func (e *Entry) Error() error {
	if strings.Index(e.Name, ErrText) != -1 || strings.Index(e.Name, ErrorText) != -1 {
		return errors.New(string(e.Content))
	}
	return nil
}
