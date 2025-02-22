// Filename: cmd/web/main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(message())
}

// Refactor to make testing easier
func message() string {
	return "Hello, world!"
}
