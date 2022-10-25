package vhost

import (
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
	"io/fs"
)

const (
	//token = byte(' ')
	eol       = byte('\n')
	comment   = "//"
	delimiter = ":"
	root      = "resource"
)

type LookupVariable = func(name string) (value string, err error)

var lookupTest = func(name string) (string, error) { return "test", nil }

var fsys fs.FS

func MountFS(f fs.FS) {
	fsys = f
}

func ReadFile(path string) ([]byte, error) {
	if path == "" {
		return nil, errors.New("invalid argument : path is empty")
	}
	if fsys == nil {
		return nil, errors.New("invalid argument : file system has not been mounted")
	}
	s, err := util.ExpandTemplate(path, lookupEnv)
	if err != nil {
		return nil, err
	}
	buf, err1 := util.FSReadFile(fsys, s)
	// If no error or there was no template, then return
	if err1 == nil || s == path {
		return buf, err1
	}
	// Override to determine if a template was used.
	s, err1 = util.ExpandTemplate(path, lookupTest)
	if err1 != nil {
		return nil, err1
	}
	return util.FSReadFile(fsys, s)
}

func ReadMap(path string) (map[string]string, error) {
	return util.FSReadMap(fsys, path)
}

var lookupEnv LookupVariable = func(name string) (string, error) {
	switch name {
	case ENV_TEMPLATE_VAR:
		return GetEnv(), nil
	}
	return "", errors.New(fmt.Sprintf("invalid argument : template variable is invalid: %v", name))
}
