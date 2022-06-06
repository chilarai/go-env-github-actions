package testing1

import (
	"log"
	"os"
)

func Test1(){
	env1, err1 := os.LookupEnv("T1")
	log.Println(env1, err1)
}