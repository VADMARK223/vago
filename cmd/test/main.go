package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "1,2,3"
	arr := strings.Split(s, ",")
	s1 := strings.Join(arr, "-")
	fmt.Println(s1)
}
