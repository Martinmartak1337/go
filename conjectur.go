package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var in int
	var o = 0
	fmt.Println("Enter a positive integer: ")
	_, err := fmt.Scanln(&in)
	if err != nil {
		os.Exit(3)
	}
	for {
		if in%2 != 0 {
			in = in*3 + 1
			o++
			fmt.Println(in)
			time.Sleep(time.Millisecond * 50)
		} else {
			in = in / 2
			o++
			if in == 4 {
				fmt.Println(o, "It's over 4 is here")
				return
			}
			fmt.Println(in)
			time.Sleep(time.Millisecond * 50)
		}
	}
}
