package main

import (
	"fmt"
	"github.com/pleycpl/godotenv"
)

func main() {
	lines := godotenv.ReadFile("./.example_env")
	for i, v := range lines {
		fmt.Println(i, " ", v)
	}
	godotenv.EnvParser(lines[0])
}
