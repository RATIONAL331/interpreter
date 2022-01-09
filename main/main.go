package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	curUser, e := user.Current()
	if e != nil {
		panic(any(e))
	}

	fmt.Printf("Hello %s! This is the Interpreter programming lanuguage!\n", curUser.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
