package main

import (
	"fmt"
	"math/rand"
)

func random() {
	ptr := new(int)
	for i := 0; i < 100; i++ {
		randomPrinter(ptr)
		fmt.Println(*ptr)
	}
}

func fooPtr(xPtr *int) {
	*xPtr = 13
}

func randomPrinter(xPtr *int) {
	x := rand.Intn(100)
	*xPtr = x
}
