package ctxreader

import (
	"context"
	"io"
	"sync"
)

func NewContextReader(ctx context.Context, r io.Reader) io.Reader {
	pr, pw := io.Pipe()

	ctxCan, cancel := context.WithCancel(ctx)

	go func() {
		_, err := io.Copy(pw, r)
		pw.CloseWithError(err)
		cancel()
	}()

	sc := &timeoutReader{
		pr:  pr,
		ctx: ctxCan,
	}

	return sc
}

type timeoutReader struct {
	once sync.Once
	pr   *io.PipeReader
	ctx  context.Context
}

func (s *timeoutReader) Read(p []byte) (n int, err error) {
	s.once.Do(func() {
		go func() {
			<-s.ctx.Done()
			s.pr.Close()
		}()
	})

	return s.pr.Read(p)
}
