package splitfs

import "context"

type reader struct {
	ctx   context.Context
	g     Getter
	pos   uint64
	stack []*node
}

func newReader(ctx context.Context, g Getter, ref []byte) (*reader, error) {
	// xxx Get the "node" at ref from g.
	return &reader{
		ctx:   ctx,
		g:     g,
		stack: []*node{n},
	}, nil
}

func (r *reader) Read(buf []byte) (int, error) {
	var n int
	for len(buf) > 0 {
		// First, unwind the stack to a node containing r.pos
		for {
			node := r.stack[len(r.stack)-1]
			if r.pos >= node.Offset && r.pos < (node.Offset+node.Size) {
				break
			}
			if len(r.stack) == 1 {
				return n, io.EOF
			}
			r.stack = r.stack[:len(r.stack)-1]
		}

		// Now walk down the tree,
		// pushing nodes onto the stack,
		// until we get to the right leaf node.
		for {
			node := r.stack[len(r.stack)-1]
			if len(node.Leaves) > 0 {
				break
			}

			index := 0
			if r.pos > node.Offset {
				// TODO: We can do better than this Search in the case where we're moving sequentially from one Node to its next sibling.
				index = sort.Search(len(node.Nodes), func(i int) bool {
					return node.Nodes[i].Offset > r.pos
				})
				index--
			}

			childRef := bs.RefFromBytes(node.Nodes[index].Ref)
			var childNode Node
			err := bs.GetProto(r.Ctx, r.g, childRef, &childNode)
			if err != nil {
				return n, errors.Wrapf(err, "getting tree node %s", childRef)
			}
			r.stack = append(r.stack, &childNode)
		}

		// Now we have a leaf node on top of the stack.
		// Discard children that precede r.pos.

		children := r.stack[len(r.stack)-1].Leaves
		for len(children) > 1 && children[1].Offset <= r.pos {
			// TODO: This could be a sort.Search
			children = children[1:]
		}

		for len(children) > 0 && len(buf) > 0 {
			offset := children[0].Offset
			ref := bs.RefFromBytes(children[0].Ref)
			chunk, err := r.g.Get(r.Ctx, ref)
			if err != nil {
				return n, errors.Wrapf(err, "getting tree node %s", ref)
			}

			// Discard any part of chunk that precedes r.pos
			skip := r.pos - offset
			chunk = chunk[skip:]
			offset += skip

			if len(chunk) >= len(buf) {
				copy(buf, chunk)
				n += len(buf)
				r.pos += uint64(len(buf))
				return n, nil
			}

			copy(buf, chunk)
			n += len(chunk)
			r.pos += uint64(len(chunk))
			buf = buf[len(chunk):]
			children = children[1:]
		}
	}
	return n, nil
}
