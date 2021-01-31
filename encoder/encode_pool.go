package encoder

import (
	"io"
	"sync"
)

var encPool = sync.Pool{
	New: func() interface{} {
		return NewEncoder(nil)
	},
}

func init() {
	for i := 0; i < 32; i++ {
		encPool.Put(NewEncoder(nil))
	}
}

// NewEncoder returns a new encoder or borrows one from the pool
func NewEncoder(w io.Writer) IEncoder {
	return &Encoder{w: w}
}

// BorrowEncoder borrows an Encoder from the pool.
func BorrowEncoder(w io.Writer) IEncoder {
	enc := encPool.Get().(*Encoder)
	enc.w = w
	enc.buf = enc.buf[:0]
	enc.isPooled = 0
	enc.err = nil
	enc.hasKeys = false
	enc.keys = nil
	return enc
}

