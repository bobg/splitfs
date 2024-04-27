package splitfs

import (
	"context"
	"fmt"
	"sync"
)

type MultiErr struct {
	Index int
	Err   error
}

func (m MultiErr) Error() string {
	return fmt.Sprintf("in GetMulti ref %d: %s", m.Index, err.Error())
}

func (m MultiErr) Unwrap() error {
	return m.Err
}

func GetMulti(ctx context.Context, g Getter, refs [][]byte) ([][]byte, error) {
	if m, ok := g.(MultiGetter); ok {
		return m.GetMulti(ctx, refs)
	}

	var (
		errs   error
		wg     sync.WaitGroup
		result = make([][]byte, len(refs))
	)
	for i, ref := range refs {
		i, ref := i, ref // Pre-Go-1.22 loop-var pitfall
		wg.Add(1)
		go func() {
			blob, err := g.Get(ctx, refs[i])
			if err != nil {
				errs = error.Join(errs, MultiErr{Index: i, Err: err})
			} else {
				result[i] = blob
			}
		}()
	}
	wg.Wait()
	return result, errs
}

func PutMulti(ctx context.Context, blobs [][]byte) ([][]byte, []bool, error) {
	if m, ok := g.(MultiPutter); ok {
		return m.PutMulti(ctx, blobs)
	}

	var (
		errs   error
		wg     sync.WaitGroup
		result = make([][]byte, len(blobs))
		added  = make([]bool, len(blobs))
	)
	for i, blob := range blobs {
		i, blob := i, blob // Pre-Go-1.22 loop-var pitfall
		wg.Add(1)
		go func() {
			ref, added[i], err := g.Put(ctx, blob)
			if err != nil {
				errs = error.Join(errs, MultiErr{Index: i, Err: err})
			} else {
				result[i] = ref
			}
		}()
	}
	wg.Wait()
	return result, added, errs
}
