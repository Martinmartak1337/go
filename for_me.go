package main

import (
	"fmt"
	"os"
)

func main() {
	const msg = "Weed?"
	var in int
	fmt.Println("Enter a positive integer: ")
	_, err := fmt.Scanln(&in)
	if err != nil {
		fmt.Printf("Error is %s", err)
		return
	}
	file, err := os.OpenFile("thc", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	for i := 0; i < in; i++ {
		_, err2 := file.WriteString(msg + " ")
		if err2 != nil {
			return
		}
	}
	err2 := file.Close()
	if err2 != nil {
		fmt.Printf("Error is %s", err2)
	}
}
