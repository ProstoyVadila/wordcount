package main

import (
	"fmt"
	"os"

	"github.com/ProstoyVadila/wordcount/count"
)

func main() {
	args := os.Args
	length := count.Count(args[1])
	fmt.Println(length)
}
