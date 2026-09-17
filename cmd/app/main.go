package main

import (
	"fmt"
	"os"
)

var version = "0.2"

func main() {
	fmt.Println("Notifier Version:", version)

	if len(os.Args) > 1 {
		fmt.Printf("Notification for %s", os.Args[1])
	} else {
		fmt.Fprintln(os.Stderr, "You must provide at least 1 argument to successfully execute the program")
		os.Exit(2)
	}

}
