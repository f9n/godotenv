package main

import (
	"fmt"
	"github.com/pleycpl/godotenv"
)

func main() {
	env := godotenv.Godotenv("./.example_env")
	fmt.Println(env)
	fmt.Println(env["MongoDb"])
}
