package main

import "fmt"

func main() {
	x := 5

	defer fmt.Println(x)

	x = 10
}
