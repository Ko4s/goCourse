package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	//Server, baza itd
	//ctrl + c
	//shutdown bazy innych i sewera

	go func() {
		for {
			time.Sleep(time.Second * 1)
			fmt.Println("Working")
		}
	}()
	done := make(chan os.Signal)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	fmt.Println("Sopting program ...")
	time.Sleep(time.Second * 1)
	fmt.Println("END")
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// wg := sync.WaitGroup{}
// wg.Add(2)
// go topic4.PrintNTimes(3, "Pierwszy", time.Second*2, &wg)
// go topic4.PrintNTimes(3, "Kolejny", time.Second*1, &wg)
// wg.Wait()
// fmt.Println("END")
