package splitfs

import (
	"io/fs"
)

type File struct {
}

var _ fs.File = &File{}

func (f *File) Stat() (fs.FileInfo, error) {
	// xxx
}

func (f *File) Read(p []byte) (n int, err error) {
	// xxx
}

func (f *File) Close() error {
	// xxx
}
