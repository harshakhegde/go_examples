package main

import (
	"fmt"
	"time"
)

func generator(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
		time.Sleep(time.Second * 3)
	}
}

func printer(c <-chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func foo() {
	c := make(chan int)
	go printer(c)
	go generator(c)

	var input string
	fmt.Scanln(&input)
}
