package main

import (
	"fmt"
	"github/Ko4s/goCourse/topic4"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// go topic4.PrintEveryNSeconds(time.Second*2, "Jestem Pierwszy", 3) // go tworzy z funkcji goroutine
	// topic4.PrintEveryNSeconds(time.Second*1, "Jestem Kolejny", 3)

	// c := make(chan bool)
	// go topic4.SetChanelE(c)
	// <-c //<- mówi ze próbujemy zwrócić wartość z chanela
	// e := <-c
	// fmt.Println(e)
	// fmt.Println("Koniec prorgamu")
	// //chanel <- true //przypisanie wartosci do channela
	// //  <- chanel zablokuje wątek, do momentu przypisania jakiejś wartości w chanelu
	// // e := <- c | do zminnej e przypisuje wartość z channelu c

	go topic4.IfiniteLoop()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
	fmt.Println("Zamykam system ....")
	time.Sleep(time.Second * 2)
	fmt.Println("END")
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
