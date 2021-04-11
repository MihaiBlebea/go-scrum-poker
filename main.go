package main

import (
	"fmt"

	"github.com/MihaiBlebea/go-scrum-poker/cmd"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	cmd.Execute()
}
