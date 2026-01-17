package main

// go mod init bondarik.net/hello
// go get golang.org/x/example/hello/reverse
// go run .
// go get golang.org/x/example/hello/reverse

// go work init ./hello

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello"), reverse.Int(12345))
}
