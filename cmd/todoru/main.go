package main

import (
	"fmt"
	"os"

	"github.com/kinbiko/todoru"
)

func main() {
	out, err := todoru.Run(os.Args[1:], os.ExpandEnv("$HOME/.todo-stack"))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(out)
}
