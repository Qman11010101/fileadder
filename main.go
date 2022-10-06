package main

import (
	"fmt"
	"os"
)

func main() {
	fp, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	fp.WriteString("TESTING")
}
