package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "	Hello World	"
	noWhiteSpace := strings.TrimSpace(str)
	fmt.Println("Before: " + str)
	fmt.Println("After: " + noWhiteSpace)
}
