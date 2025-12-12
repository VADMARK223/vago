package main

import "fmt"

func main() {
	fmt.Println(f())
}

func f() (x int) {
	defer func() {
		x++
	}()
	return 1
}
