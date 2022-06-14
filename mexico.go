package main

import (
	"fmt"
	"github.com/thanhpk/randstr"
	"os"
)

func main() {
	e, jjj := os.Executable()
	if jjj != nil {
		return
	}
	err420 := os.Rename(e, "Sup"+randstr.Base64(16)+"Martin")
	if err420 != nil {
		return
	}
	var in string
	fmt.Println("Enter a string!")
	_, uuu := fmt.Scanln(&in)
	if uuu != nil {
		return
	}
	fmt.Println(acceptString(in))
}

func acceptString(str string) string {
	byteStr := []rune(str)
	for i, j := 0, len(byteStr)-1; i < j; i, j = i+1, j-1 {
		byteStr[i], byteStr[j] = byteStr[j], byteStr[i]
	}
	return string(byteStr)
}

// Idk it's stupid
