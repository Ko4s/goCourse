package topic4

import (
	"fmt"
	"sync"
	"time"
)

//PrintNTimes prints msg every t second, n times
func PrintNTimes(n int, msg string, t time.Duration, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		fmt.Println(msg)
		time.Sleep(t)
	}
	wg.Done()
}

func PrintNTimesC(n int, msg string, t time.Duration, c chan<- int) {
	for i := 0; i < n; i++ {
		fmt.Println(msg)
		time.Sleep(t)
	}
	c <- n
	//close()
}

func Increment(n *int, x int, wg *sync.WaitGroup) {

	for i := 0; i < x; i++ {
		*n++
	}
	fmt.Println(*n)
	wg.Done()
}
