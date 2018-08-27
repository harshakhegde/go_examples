package main

import (
	"fmt"
	"math/rand"
	"time"
)

func trd(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func es() {
	for i := 0; i < 10; i++ {
		go trd(i)
	}
	var res int
	fmt.Scanln(&res)
}
