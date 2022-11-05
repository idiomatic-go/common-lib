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

func ProcessContent[T any](ctx context.Context) (T, error) {
	var t T
	if ctx == nil {
		return t, errors.New(fmt.Sprintf("invalid argument : context is nil"))
	}
	fs := ContextContent(ctx)
	if fs == nil {
		return t, errors.New(fmt.Sprintf("no file system entry available"))
	}
	if fs.Content == nil || len(fs.Content) == 0 {
		return t, errors.New(fmt.Sprintf("no content available for entry name : %v", fs.Name))
	}
	err := fs.Error()
	if err != nil {
		return t, err
	}
	if strings.Index(fs.Name, ".json") == -1 {
		return t, errors.New(fmt.Sprintf("invalid content for json.Unmarshal() : %v", fs.Name))
	}
	err = json.Unmarshal(fs.Content, &t)
	return t, err
}
