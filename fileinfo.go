package splitfs

import (
	"io/fs"
	"time"
)

type FileInfo struct {
}

var _ fs.FileInfo = &FileInfo{}

func (f *FileInfo) Name() string {
	// xxx
}

func (f *FileInfo) Size() int64 {
	// xxx
}

func (f *FileInfo) Mode() fs.FileMode {
	// xxx
}

func (f *FileInfo) ModTime() time.Time {
	// xxx
}

func (f *FileInfo) IsDir() bool {
	// xxx
}

func (f *FileInfo) Sys() any {
	// xxx
}
