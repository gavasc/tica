package main

import (
	"fmt"
	"os"
)

func main() {
  args := os.Args
  command := args[1]

  switch command {
    case "punch":
      if len(args) >= 3 {
        punchHandler(args[2:])
      } else {
        fmt.Println("Missing task code!")
      }
  }
}

