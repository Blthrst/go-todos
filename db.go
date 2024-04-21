package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func InitSecrets() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
}