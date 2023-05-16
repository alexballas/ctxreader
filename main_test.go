package ctxreader

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"
)

func TestNewContextReader(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	reader := NewContextReader(ctx, strings.NewReader("Test"))
	cancel()

	_, err := io.Copy(io.Discard, reader)

	if !errors.Is(err, context.Canceled) {
		t.Errorf("Unexpected error. Expected '%s', got '%s'", context.Canceled, err)
	}
}
