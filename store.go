package splitfs

// Getter can read content-addressable data blobs.
type Getter interface {
	Get(context.Context, []byte) ([]byte, error)
	List(context.Context, func([]byte) error) error
}

// MultiGetter is an interface that Getters may optionally implement to make the GetMulti function efficient.
type MultiGetter interface {
	GetMulti(context.Context, [][]byte) ([][]byte, error)
}

type Store interface {
	Getter

	// Put adds a blob to the store if it was not already present.
	// It returns the blob's ref and a boolean that is true iff the blob was added (vs already present).
	Put(context.Context, []byte) (ref []byte, added bool, err error)
}

type MultiPutter interface {
	PutMulti(context.Context, [][]byte) ([][]byte, []bool, error)
}
