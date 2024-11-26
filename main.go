package main

import (
	"fmt"
	Application "gowb/app"
	Server "gowb/server"
	"os"
)

func main() {

	app := Application.Application{}

	args := os.Args
	arg_len := len(args)

	for i := 1; i < arg_len; i++ {
		switch args[i] {
		case "-h":
			fmt.Println("User Asking for help")
		case "new":
			if i == arg_len-1 { // means this is the last argument use provided and user did not pass the application name
				fmt.Println("Please Provide a application name")
				os.Exit(-1)
			}
			i++
			app.New(args[i])
		case "serve":
			server := Server.Server{}
			server.RunServer()
		}
	}
}
