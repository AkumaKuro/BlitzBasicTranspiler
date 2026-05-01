package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, _ := os.ReadFile("./examples/example001.bb")
	source := string(bytes)

	fmt.Printf("%s\n", source)
}
