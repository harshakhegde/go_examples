package main

import (
	"fmt"
)

func fooStruct() {
	d := new(Dog)
	d.leg = 4

	d.sound = "wooofff"
	s := d.bark()
	fmt.Println(s)
	fmt.Println(d.leg)
	fmt.Println(d.hello())
}

//Dog //
type Dog struct {
	leg   int
	sound string
}
//Animal //
type Animal interface {
	hello() string
}

func (d *Dog) bark() string {
	return d.sound
}

func (d *Dog) hello() string {
	return d.sound
}

func bark(a Animal) string {
	return a.hello()
}
