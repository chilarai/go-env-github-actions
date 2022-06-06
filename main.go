package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main(){
	fmt.Println("Go env github actions")

	env1, err1 := os.LookupEnv("T1")
	log.Println(env1, err1)

	fmt.Println(os.Environ())

	if !err1 {
		fmt.Println("Could not read TEST1 from env")
	} else {
		fmt.Println("Value of Test1", env1)
	}
}