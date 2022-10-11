package vhost

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
	"io"
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
	buf, err1 := fs.ReadFile(fsys, root+"/"+s)
	// If no error or there was no template, then return
	if err1 == nil || s == path {
		return buf, err1
	}
	// Override to determine if a template was used.
	s, err1 = util.ExpandTemplate(path, lookupTest)
	if err1 != nil {
		return nil, err1
	}
	return fs.ReadFile(fsys, root+"/"+s)
}

func ReadMap(path string) (map[string]string, error) {
	buf, err := ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParseBuffer(buf)
}

func ParseBuffer(buf []byte) (map[string]string, error) {
	m := make(map[string]string)
	if len(buf) == 0 {
		return m, nil
	}
	r := bytes.NewReader(buf)
	reader := bufio.NewReader(r)
	var line string
	var err error
	for {
		line, err = reader.ReadString('\n')
		k, v, err0 := ParseLine(line)
		if err0 != nil {
			return m, err0
		}
		if len(k) > 0 {
			m[k] = v
		}
		if err == io.EOF {
			break
		} else {
			if err != nil {
				break
			}
		}
	}
	return m, nil
}

func ParseLine(line string) (string, string, error) {
	if len(line) == 0 {
		return "", "", nil
	}
	line = strings.TrimLeft(line, " ")
	if len(line) == 0 || strings.HasPrefix(line, comment) {
		return "", "", nil
	}
	i := strings.Index(line, delimiter)
	if i == -1 {
		return "", "", fmt.Errorf("invalid argument : line does not contain the ':' delimeter : [%v]", line)
	}
	key := line[:i]
	val := line[i+1:]
	return strings.TrimSpace(key), strings.TrimLeft(val, " "), nil
}

var lookupEnv LookupVariable = func(name string) (string, error) {
	switch name {
	case ENV_TEMPLATE_VAR:
		return GetEnv(), nil
	}
	return "", errors.New(fmt.Sprintf("invalid argument : template variable is invalid: %v", name))
}
