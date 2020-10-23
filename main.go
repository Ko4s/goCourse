package main

import (
	"github/Ko4s/goCourse/topic4"
	"log"
	"time"
)

func main() {
	// Go routines example
	// go topic4.InfinitePrintFunction("Bump") //odpalenie gorutyny
	// time.Sleep(time.Second) //czekam 1s, cały program kończy się kiedy skonczy się funkcja main

	//Channele
	c := make(chan int)
	go topic4.SetChanel(c)
	go topic4.GetChannelValue(c)
	time.Sleep(time.Second * 10)
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
