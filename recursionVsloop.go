package main

import (
	"fmt"
	"time"
)

func main() {
	compareRecAndLoop(10000);
}




func compareRecAndLoop(ct int) {
	fmt.Println("Cal fibby for", ct)
	start1 := time.Now()
	fibbyRec(ct, 0, 1)
	t1 := time.Now()
	elapsed1 := t1.Sub(start1)
	fmt.Println("using rec", elapsed1)

	start2 := time.Now()
	fibby(ct)
	t2 := time.Now()
	elapsed2 := t2.Sub(start2)
	fmt.Println("using loop", elapsed2)
}

func fibby(ct int) {
	var x  = 0
	var y  = 1
	var temp = x
	var prev int
	for i := 0; i < ct; i++ {
		temp = temp + prev
		//fmt.Println(i, temp)
		prev = y
		y = temp
	}
}

func fibbyRec(count int, cur int, nxt int) {
	if count == 0 {
		return
	}
	//fmt.Println(cur)
	fibbyRec(count-1, nxt, cur+nxt)
}

