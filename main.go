package main

import (
	"fmt"
	"github.com/limyel/myinterpreter/repl"
	"os"
	"os/user"
)

func main() {
	crtUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is MyInterpreter!\n", crtUser.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
