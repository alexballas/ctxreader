package ctxreader

import (
	"context"
	"io"
	"sync"
)

// NewContextReader creates a new io.Reader with a context.
func NewContextReader(ctx context.Context, r io.Reader) io.Reader {
	if ctx == nil {
		ctx = context.Background()
	}

	if r == nil {
		return nil
	}

	pr, pw := io.Pipe()

	ctxCan, cancel := context.WithCancel(ctx)

	go func() {
		_, err := io.Copy(pw, r)
		pw.CloseWithError(err)
		cancel()
	}()

	cr := &ctxReader{
		pr:  pr,
		pw:  pw,
		ctx: ctxCan,
	}

	return cr
}

type ctxReader struct {
	once sync.Once
	pr   *io.PipeReader
	pw   *io.PipeWriter
	ctx  context.Context
}

func (s *ctxReader) Read(p []byte) (n int, err error) {
	s.once.Do(func() {
		go func() {
			<-s.ctx.Done()
			s.pw.CloseWithError(s.ctx.Err())
		}()
	})

	return s.pr.Read(p)
}
