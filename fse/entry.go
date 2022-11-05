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

func ProcessContent(ctx context.Context, t any) error {
	if ctx == nil {
		return errors.New(fmt.Sprintf("invalid argument : context is nil"))
	}
	fs := ContextContent(ctx)
	if fs == nil {
		return errors.New(fmt.Sprintf("no file system entry available"))
	}
	if fs.Content == nil || len(fs.Content) == 0 {
		return errors.New(fmt.Sprintf("no content available for entry name : %v", fs.Name))
	}
	err := fs.Error()
	if err != nil {
		return err
	}
	if t == nil {
		return errors.New("invalid argument : any parameter is nil")
	}
	if strings.Index(fs.Name, ".json") == -1 {
		return errors.New(fmt.Sprintf("invalid content for json.Unmarshal() : %v", fs.Name))
	}
	return json.Unmarshal(fs.Content, t)
}
