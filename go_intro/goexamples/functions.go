package goexamples

import (
	"fmt"
	"log"
)

func DestroyWorld() {
	fmt.Print("Destroy!!")
	imie := "Piotrek"
	pointer(&imie) //& oznacza ze chcem pointer
	isloged, err := login("abc", "12345")
	if err != nil {
		log.Fatalf("")
	}

	fmt.Println(isloged)
}

func pointer(str *string) {
	fmt.Println(*str)
}

func login(userName, password string) (bool, error) {
	return true, nil
}
