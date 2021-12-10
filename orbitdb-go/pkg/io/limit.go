package io

import (
	"fmt"
)

type ErrLimitReader struct {
	l int64
}

func (e ErrLimitReader) Error() string {
	return fmt.Sprintf("hit limit reader limit %d", e.l)
}

//

func NewErrLimitReader(l int64) ErrLimitReader { return ErrLimitReader{l} }

//

type LimitReader struct {
	r   Reader // underlying reader
	l   int64  // initial limit
	n   int64  // max bytes remaining
	err error  // sticky error
}

func (l *LimitReader) Read(p []byte) (n int, err error) {
	if l.err != nil {
		return 0, l.err
	}
	if len(p) == 0 {
		return 0, nil
	}
	// If they asked for a 32KB byte read but only 5 bytes are
	// remaining, no need to read 32KB. 6 bytes will answer the
	// question of the whether we hit the limit or go past it.
	if int64(len(p)) > l.n+1 {
		p = p[:l.n+1]
	}
	n, err = l.r.Read(p)

	if int64(n) <= l.n {
		l.n -= int64(n)
		l.err = err
		return n, err
	}

	n = int(l.n)
	l.n = 0
	l.err = NewErrLimitReader(l.l)
	return n, l.err
}

//

// LimitReader is similar to io.LimitReader but returns
// constant error in case we have hit the limit.
//
func NewLimitReader(r Reader, l int64) Reader {
	return &LimitReader{r: r, l: l, n: l}
}
