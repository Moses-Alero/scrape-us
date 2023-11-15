package main

import (
	"fmt"
	"log"

	"scrape-us/cmd"
)

type error interface {
	Error() string
}

func main() {
	fmt.Println("Hello, World!")
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}

}
