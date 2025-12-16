package main

import "fmt"

func tryRecover() {
	if err := recover(); err != nil {
		fmt.Println("covered")
	}
}

func main() {
	defer fmt.Println("defer")
	defer func() {
		fmt.Println("end")
		tryRecover()
	}()
	fmt.Println("start")
	panic("panic")
}

// start end covered defer
// start end defer
// start defer covered end
// start covered end

// start defer end covered
// start defer covered end
// start covered end defer
