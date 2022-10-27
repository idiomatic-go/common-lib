package vhost

import (
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/fse"
	"github.com/idiomatic-go/common-lib/util"
	"io/fs"
)

const (
	root = "resource"
)

type LookupVariable = func(name string) (value string, err error)

var lookupTest = func(name string) (string, error) { return "test", nil }

var fsys fs.FS

func MountFS(f fs.FS) {
	fsys = f
}

// ReadFile - read a file from the mounted fs, adding the resource directory as that is not known by the client
func ReadFile(name string) ([]byte, error) {
	if name == "" {
		return nil, fmt.Errorf("invalid argument : file name is empty")
	}
	if fsys == nil {
		return nil, errors.New("invalid argument : file system has not mounted")
	}
	s, err := util.ExpandTemplate(name, lookupEnv)
	if err != nil {
		return nil, err
	}
	buf, err1 := fse.ReadFile(fsys, root+"/"+s)
	// If no error or there was no template, then return
	if err1 == nil || s == name {
		return buf, err1
	}
	// Override to determine if a template was used.
	s, err1 = util.ExpandTemplate(name, lookupTest)
	if err1 != nil {
		return nil, err1
	}
	return fse.ReadFile(fsys, root+"/"+s)
}

func ReadMap(path string) (map[string]string, error) {
	return fse.ReadMap(fsys, path)
}

var lookupEnv LookupVariable = func(name string) (string, error) {
	switch name {
	case EnvTemplateVar:
		return GetEnv(), nil
	}
	return "", errors.New(fmt.Sprintf("invalid argument : template variable is invalid: %v", name))
}
