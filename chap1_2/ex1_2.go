// Ex 1.2 prints its command-line arguments.
// with index and value, one per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}
}
