package topic4

import (
	"fmt"
	"time"
)

/**
For whole example look into main function
*/
func InfinitePrintFunction(msg string) {
	for {
		fmt.Println(msg)
		time.Sleep(time.Millisecond)
	}
}

func changeBoolChanel(c chan bool) {
	c <- true //im seting c channel for true
}

func SetChanel(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	//close(c)
	fmt.Println("Seting done")
}

func GetChannelValue(c chan int) {

	for v := range c {
		//time.Sleep(time.Millisecond)
		fmt.Println(v)
	}
	fmt.Println("RECEIVER DONE")
}
