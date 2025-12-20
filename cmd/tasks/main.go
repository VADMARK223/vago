package main

import (
	"fmt"
	"vago/cmd/test"
)

func main() {
	fmt.Printf("➡️ \033[93m%s: \033[92m%v\033[0m\n", "sum:", test.Abs(0))
}
