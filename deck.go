package main

import "fmt"

type deck []string

func (d deck) print() {
	fmt.Println(d)
}

func (d deck) getTitle() deck {
	return d
}
