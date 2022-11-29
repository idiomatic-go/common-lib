package fse

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func (e *Entry) Error() error {
	if strings.Index(e.Name, ErrText) != -1 || strings.Index(e.Name, ErrorText) != -1 {
		return errors.New(string(e.Content))
	}
	return nil
}

// ProcessContent - process content
func ProcessContent[T any](ctx context.Context) (T, error) {
	var t T
	if ctx == nil {
		return t, errors.New(fmt.Sprintf("fse:ProcessContent internal error : context is nil"))
	}
	return ProcessEntry[T](ContextContent(ctx))
}

// ProcessEntry - process content that is fs entry
func ProcessEntry[T any](fs *Entry) (T, error) {
	var t T

	if fs == nil {
		return t, errors.New(fmt.Sprintf("fse:ProcessEntry internal error : no file system entry available"))
	}
	if fs.Content == nil || len(fs.Content) == 0 {
		return t, errors.New(fmt.Sprintf("fse:ProcessEntry internal error : no content available for entry name : %v", fs.Name))
	}
	err := fs.Error()
	if err != nil {
		return t, err
	}
	if strings.Index(fs.Name, ".json") == -1 {
		return t, errors.New(fmt.Sprintf("fse:ProcessEntry internal error : invalid content for json.Unmarshal() : %v", fs.Name))
	}
	err = json.Unmarshal(fs.Content, &t)
	if err != nil {
		return t, errors.New(fmt.Sprintf("fse:ProcessEntry internal error : json.Unmarshal() : %v", err))
	}
	return t, nil
}
