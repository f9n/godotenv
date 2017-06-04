package main

import (
	"fmt"
	"github.com/pleycpl/godotenv"
)

func main() {
	lines := godotenv.ReadFile("../godotenv.go")
	for i, v := range lines {
		fmt.Println(i, " ", v)
	}
}
