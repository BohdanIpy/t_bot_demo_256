package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("bot started")

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Getenv("TOKEN"))
}
