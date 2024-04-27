package splitfs

import "io/fs"

type DirEntry struct {
}

var _ fs.DirEntry = &DirEntry{}

func (d *DirEntry) Info() (fs.FileInfo, error) {
	// xxx
}

func (d *DirEntry) IsDir() bool {
	// xxx
}

func (d *DirEntry) Name() string {
	// xxx
}

func (d *DirEntry) Type() fs.FileMode {
	// xxx
}
