package vhost

import (
	"strings"
)

type errorList struct {
	errs []error
}

func (e *errorList) Error() string {
	if !e.IsError() {
		return ""
	}
	return e.Cat()
}

func (e *errorList) IsError() bool { return len(e.errs) != 0 }

func (e *errorList) Errors() []error {
	return e.errs
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

//func (e *errorList) Handled() {
//	e.errs = nil
//}

func NewErrors(errs ...error) Errors {
	return NewErrorsList(errs)
}

func NewErrorsList(errs []error) Errors {
	s := errorList{}
	for _, e := range errs {
		if e == nil {
			continue
		}
		s.errs = append(s.errs, e)
	}
	return &s
}

func NewErrorsAny(a any) Errors {
	s := errorList{}
	if IsNil(a) {
		return &s
	}
	if err, ok := a.(error); ok {
		s.errs = append(s.errs, err)
	}
	return &s
}

func newErrors() Errors {
	return &errorList{}
}
