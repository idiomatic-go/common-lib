package vhost

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"strings"
)

const (
	//token = byte(' ')
	eol     = byte('\n')
	comment = "//"
	space   = " "
)

var fsys fs.FS
var dir = "resource"

func MountFS(f fs.FS, directory string) {
	fsys = f
	if len(directory) > 0 {
		dir = directory
	}
}

func ReadFile(name string) ([]byte, error) {
	if len(name) == 0 {
		return nil, errors.New("invalid argument : file name is empty")
	}
	if fsys == nil {
		return nil, errors.New("invalid argument : file system has not been mounted")
	}
	return fs.ReadFile(fsys, dir+"/"+name)
}

func ReadMap(name string) (map[string]string, error) {
	buf, err := ReadFile(name)
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
	buffer := bytes.NewBuffer(buf)
	var line string
	var err error
	for line, err = buffer.ReadString(eol); err != nil; {
		k, v, err0 := ParseLine(line)
		if err0 != nil {
			return m, err0
		}
		if len(k) > 0 {
			m[k] = v
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
	i := strings.Index(line, space)
	if i == -1 {
		return "", "", fmt.Errorf("invalid argument : line does not contain the space ' ' delimeter : [%v]", line)
	}
	//key := line[:i]
	//val := line[i+1:]
	//m[strings.TrimSpace(key)] = strings.TrimLeft(val, " ")
	return strings.TrimSpace(line[:i]), strings.TrimLeft(line[i:], " "), nil
}
