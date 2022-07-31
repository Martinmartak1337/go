package main

import (
	"fmt"
	"time"
)

func main() {
	var in int
	var start, nanoS int64
	var i int
	fmt.Println("Enter an int!")
	_, err420 := fmt.Scanln(&in)
	if err420 != nil {
		return
	}
	start = time.Now().UnixNano()
	for {
		if i >= in {
			break
		}
		count(4, 6)
		i++
	}
	nanoS = time.Now().UnixNano() - start
	fmt.Println(nanoS)
}

func count(a int64, b int64) int64 {
	return a + b
}
