# Context-aware io.Reader [![Go Reference](https://pkg.go.dev/badge/github.com/alexballas/ctxreader.svg)](https://pkg.go.dev/github.com/alexballas/ctxreader)

## Example:
```go
package main

import (
	"bufio"
	"context"
	"ctxreader"
	"fmt"
	"os"
	"time"
)

// This example shows how to use NewContextReader.
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// create the context reader.
	cr := ctxreader.NewContextReader(ctx, os.Stdin)

	// read our new reader.
	buf := bufio.NewScanner(cr)
	fmt.Println("You have 5 seconds to enter your input:")
	buf.Scan()

	switch buf.Err() {
	case nil:
		fmt.Println("You entered:", buf.Text())
	default:
		fmt.Println("Time is up!", buf.Err())
	}
}
```