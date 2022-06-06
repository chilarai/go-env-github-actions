package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T){
	main()

	os.Setenv("ENV_VAR", "value")
    defer os.Unsetenv("ENV_VAR")
}