package fse

const (
	// Scheme - file system entry
	Scheme = "fse"
)

type Entry struct {
	Name    string
	Content []byte
}
