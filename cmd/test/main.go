package main

import "fmt"

func f(p *int) {
	p = nil
}

func main() {
	a := 10
	p := &a
	f(p)
	fmt.Println(*p)
}
