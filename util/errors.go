package util

import (
	"strings"
)

type errorList struct {
	errs []error
}

func (e *errorList) Errors() []error {
	return e.errs
}

func (e *errorList) Error() string {
	if len(e.errs) == 0 {
		return ""
	}
	return e.errs[0].Error()
}

func (e *errorList) Add(err error) {
	if err != nil {
		e.errs = append(e.errs, err)
	}
}

func (e *errorList) Cat() string {
	if len(e.errs) == 0 {
		return ""
	}
	//sep := " "
	var sb strings.Builder
	//sb.WriteString(fmt.Sprintf("%v:%v", sc.first, sc.errs[sc.first].Error()))
	for i, err := range e.errs {
		sb.WriteString(err.Error())
		if i < len(e.errs)-1 {
			sb.WriteString(" : ")
		}
	}
	return sb.String()
}

func NewErrors(errs ...error) Errors {
	s := errorList{}
	for _, e := range errs {
		if e == nil {
			continue
		}
		s.errs = append(s.errs, e)
	}
	return &s
}

func newErrors() Errors {
	return &errorList{}
}
