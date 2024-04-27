package splitfs

import "io/fs"

type FS struct {
}

var (
	_ fs.FS        = &FS{}
	_ fs.ReadDirFS = &FS{}
	_ fs.StatFS    = &FS{}
)

func (f *FS) Open(name string) (fs.File, error) {
	// xxx
}

func (f *FS) ReadDir(name string) ([]fs.DirEntry, error) {
	// xxx
}

func (f *FS) Stat(name string) (fs.FileInfo, error) {
	// xxx
}
