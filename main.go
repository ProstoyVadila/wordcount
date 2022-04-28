package main

import (
	"fmt"
	"os"

	"github.com/ProstoyVadila/wordcount/count"
)

func main() {
	var length int
	args := os.Args

	if len(args) > 1 {
		length = count.Count(args[1])
	} else {
		length = 0
	}
	fmt.Println(length)
}
