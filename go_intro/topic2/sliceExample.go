package topic2

import "fmt"

var exampleVariable string
var arrayExample = [7]int{1, 2, 3, 4}
var sliceExample = []int{}

func ShowExamples() {
	arrayExample[6] = 7
	//dodawnie elementu do slice
	fmt.Println(sliceExample)
	sliceExample = append(sliceExample, 1, 2, 3, 4)
	//fmt.Println(arrayExample)
	fmt.Println(sliceExample)
	//remove element od index 2
	sliceExample = append(sliceExample[:2], sliceExample[3:]...)
	fmt.Println(sliceExample)
	str := "Piotrek"
	str = string(append([]byte(str), byte(77)))
	fmt.Println(str)
}
