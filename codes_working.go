package main

import "fmt"

func main() {
	var even int64
	var odd int64
	even, odd = c(940054700987)
	fmt.Println(even, odd)
}

func c(brcd int64) (int64, int64) {
	var even int64
	var odd int64
	for {
		if brcd != 0 {
			odd = odd + (brcd % 10)
			brcd = brcd / 10
		} else {
			return even, odd
		}
		if brcd != 0 {
			even = even + (brcd % 10)
			brcd = brcd / 10
		} else {
			return even, odd
		}
	}
}
