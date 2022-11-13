package util

type errorList struct {
	errs []error
}

//S struct {
//	errs []error
//}

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
