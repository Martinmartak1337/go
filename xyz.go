package main

import (
	"fmt"
	"os"
)

func main() {
	s(90, 120, 0.5)
}

func s(v1 float32, v2 float32, dt float32) {
	if v1 >= v2 {
		fmt.Println("v2 must be greater than v1!")
		os.Exit(3)
	}
	if dt < 0 {
		fmt.Println("dt must be positive!")
		os.Exit(3)
	}
	t2 := (v1 * dt) / (v2 - v1)
	s2 := v2 * t2
	fmt.Println(s2)
}
