package main

import "vago/internal/app"

func main() {
	s := []int{1, 2, 3}
	s2 := s[:1]
	s3 := append(s2, 4)
	app.Dump("", s)
	app.Dump("", s2)
	app.Dump("", s3)

	app.Dump("le", len(s))
	app.Dump("le", cap(s))

	/*s := []int{10, 20, 30}
	s2 := s[:2]
	app.Dump("", s2)
	app.Dump("", len(s2))
	app.Dump("", cap(s2))
	s2[1] = 99
	app.Dump("", s[1])*/

	/*log.Println("start")
	ch := make(chan int)
	go func() {
		ch <- 10
	}()
	log.Println(<-ch)*/

	/*s := []int{1, 2, 3}
	app.Dump("len", len(s))
	app.Dump("cap", cap(s))
	app.Dump("s", s)

	s2 := append(s, 4)
	app.Dump("len", len(s2))
	app.Dump("cap", cap(s2))
	app.Dump("s2", s2)
	s2[0] = 99
	fmt.Println(s[0])*/

	/*s := []int{1, 2, 3}
	s2 := append(s, 4)
	s2[0] = 99
	fmt.Println(s[0])*/

}
