package httpxt

import "io/fs"

var fsys fs.FS

func Mount(f fs.FS) {
	fsys = f
}

func UnmountFS() {
	fsys = nil
}
