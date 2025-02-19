package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("Missing command!")
		return
	}
	command := args[1]

	switch command {
	case "punch":
		if len(args) >= 3 {
			PunchHandler(args[2:]) // TODO change to pass only the task code string
		} else {
			fmt.Println("Missing task code!")
		}
	case "list":
		fmt.Println(ListHandler())
	case "delete":
		if len(args) >= 3 {
			DeleteHandler(args[2])
		} else {
			fmt.Println("Missing task code!")
		}
	default:
		fmt.Println("Unknown command")
	}
}
