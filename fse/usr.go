package fse

const (
	// Scheme - file system entry
	Scheme    = "fse"
	ErrorText = "error.txt"
	ErrText   = "err.txt"
)

type Entry struct {
	Name    string
	Content []byte
}
