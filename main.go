package main

import (
	"embed"
	"fmt"

	"github.com/MihaiBlebea/go-scrum-poker/cmd"
	"github.com/joho/godotenv"
)

//go:embed webapp/dist/*
var static embed.FS

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	cmd.Execute()
}
