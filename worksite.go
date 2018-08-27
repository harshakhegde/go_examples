package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mainy() {
	compareRecAndLoop(10000)
}
func calc() {
	var a int
	var b int
	fmt.Println("Enter first number")
	fmt.Scanln(&a)
	fmt.Println("Enter second number")
	fmt.Scanln(&b)

	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Second)

	fmt.Println("Sum is ", a+b)

}

func loop() {
	for i := 0; i < 100; i++ {
		var x string
		if i%3 == 0 {
			x = "Bingo"
		}
		if i%5 == 0 {
			x = x + "Tringo"
		}
		fmt.Println(i, x)
	}
}

func slicing() {
	arr := [5]int{1, 2, 3, 4, 5}
	x := arr[2:]
	src := arr[:]

	y := append(x, 6)
	fmt.Println(y)
	trg := make([]int, 4)
	copy(trg, src)
	fmt.Println(trg)
}

func mapy() {
	m := make(map[int]int)
	for i := 0; i < 100; i++ {
		m[i] = i * i
	}
	size := len(m)
	fmt.Println(size)
	fmt.Println(m[90])
	fmt.Println(m[size-1])
}

func otherFor() {
	arr := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}

	for _, value := range arr {
		fmt.Println(value)
	}
}

func fooq() {
	x, y := returnTwin()
	fmt.Println(x, y)
}

func returnTwin() (string, string) {
	return "sum", "sim"
}

func variadicFunc() {
	res := summer(1, 2, 3, 4, 5)
	fmt.Println(res)
}

func summer(x ...int) int {
	var sum int
	for _, value := range x {
		sum += value
	}
	return sum
}

func closure() {
	x := func() func() int {
		return func() int {
			return 1
		}
	}

	fmt.Println(x()())
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
