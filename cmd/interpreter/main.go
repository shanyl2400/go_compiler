package main

import (
	"fmt"
	"os"
	"os/user"
	"shanyl2400/go_compiler/repl"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, This is go_interpreter programing language!\n", usr.Username)
	fmt.Printf("Feel free to type commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
