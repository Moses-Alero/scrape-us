package main

import (
	"fmt"
	"log"
	"time"

	"scrape-us/cmd"
)

type error interface {
	Error() string
}

func main() {
	fmt.Println("I'm gonna Scrape you so good :p ")
	time.Sleep(time.Second)
	go Spinner()
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}

}

func Spinner() {
	spinner := `-\|/`
	for i := 0; i < 200000; i++ {
		fmt.Printf("\rScraping... %c", spinner[i%len(spinner)])
		time.Sleep(100 * time.Millisecond)
	}
}
