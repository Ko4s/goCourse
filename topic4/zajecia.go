package topic4

import (
	"fmt"
	"time"
)

//PrintEveryNSeconds msg every n seconds
func PrintEveryNSeconds(t time.Duration, msg string, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(msg, ":", t)
		time.Sleep(t)
	}
}

func SetChanelE(c chan bool) {
	time.Sleep(time.Second * 7)
	c <- true //do channelu dodaje wartość true
	c <- false
}

func IfiniteLoop() {
	for {
		time.Sleep(time.Second)
	}
}
