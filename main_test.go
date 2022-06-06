package main

import (
	"log"
	"os"
	"testing"
)

func TestMain(t *testing.T){
	main()

	env1, err1 := os.LookupEnv("T1")
	log.Println(env1, err1)
}