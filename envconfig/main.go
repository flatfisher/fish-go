package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	Gopath string
}

func main() {
	var goenv Env
	envconfig.Process("", &goenv)
	fmt.Println(goenv)
}
