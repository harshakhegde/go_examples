package main

import (
	"fmt"
)

func counter(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
}

func isEven(c chan int) {
	for {
		i := <-c
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println("isEven: I got", i, " but I cannot process it")
			// r := rand.Intn(10)
			// time.Sleep(time.Duration(r) * time.Second)
			// c <- i
		}
	}
}

//Channel sniffs for odd numbers
func isOdd(c chan int) {
	for {
		i := <-c
		if i%2 != 0 {
			fmt.Println(i, "is odd")
		} else {
			fmt.Println("isOdd: I got", i, " but I cannot process it")
			// r := rand.Intn(10)
			// time.Sleep(time.Duration(r) * time.Second)
			// c <- i
		}
	}
}

func call() {
	c := make(chan int)
	go counter(c)
	go isEven(c)
	go isOdd(c)

	var s string
	fmt.Scanln(&s)
}
