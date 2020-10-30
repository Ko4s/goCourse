package topic2

import "fmt"

//Przykład array -> [ilosc_elementow]typ_tablicy{}
var arrayExample = [7]int{1, 2, 3, 4, 5, 6}
var sliceExample = []int{10, 12, 13, 14, 15, 16}
var arrayMakeExample = make([]string, 0)

//ShowExamples show examples
func ShowExamples() {
	arrayExample[6] = 10
	sliceExample = append(sliceExample, 1, 2, 3, 4, 5) //tak dodaje się elementy w go do slice'a
	sliceExample = append(sliceExample[:2], sliceExample[3:]...)
	fmt.Println(arrayExample[:3])
	fmt.Println(arrayExample) //
	fmt.Println(sliceExample)
}
