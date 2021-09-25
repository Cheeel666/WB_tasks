package main

import (
	"fmt"
	"os"
	wget "task9/wget"
)

// Usage: go run main.go wget test.txt https://docs.google.com/document/d/1R7dbON52NlBLYun8crm4nHtedS1wZ9TCdQrWRLjIEio/edit
func main() {
	args := os.Args[1:]
	switch args[0] {
	case "wget":
		err := wget.Wget(args[1:])
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("Incorrect command")
	}
}
