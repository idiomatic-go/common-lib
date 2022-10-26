package fse

import (
	"context"
	"fmt"
	"io/fs"
	"strings"
)

const (
	//token = byte(' ')
	eol       = byte('\n')
	comment   = "//"
	delimiter = ":"
	root      = "resource"
)

func ReadFile(fsys fs.FS, name string) ([]byte, error) {
	if name == "" {
		return nil, fmt.Errorf("invalid argument : file name is empty")
	}
	if fsys == nil {
		return nil, fmt.Errorf("invalid argument : file system is nil")
	}
	return fs.ReadFile(fsys, name)
}

func ReadFileContext(ctx context.Context, path string) ([]byte, error) {
	if ctx == nil {
		return nil, fmt.Errorf("invalid argument : context is nil")
	}
	fsys := ContextEmbeddedFS(ctx)
	name := ContextEmbeddedContent(ctx)
	if name == "" {
		path = strings.TrimPrefix(path, "/")
		return ReadFile(fsys, path)
	}
	return ReadFile(fsys, name)
}

func ReadDir(fsys fs.FS, name string) ([]fs.DirEntry, error) {
	if fsys == nil {
		return nil, fmt.Errorf("invalid argument : file system is nil")
	}
	if name == "" {
		return nil, fmt.Errorf("invalid argument : directory name is empty")
	}
	return fs.ReadDir(fsys, name)
}

func ReadMap(fsys fs.FS, name string) (map[string]string, error) {
	buf, err := ReadFile(fsys, name)
	if err != nil {
		return nil, err
	}
	return ParseBuffer(buf)
}
