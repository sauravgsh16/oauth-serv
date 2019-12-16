package main

import (
	"github.com/sauravgsh16/oauth-serv/cmd"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		panic(err)
	}
}
