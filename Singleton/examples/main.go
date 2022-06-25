package main

import "fmt"

var singleton = &Singleton{}

type Singleton struct{}

func (s *Singleton) echo() {
	fmt.Printf("Hello!!")
}

func main() {
	singleton.echo()
}
