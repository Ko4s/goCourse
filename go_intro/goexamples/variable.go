package goexamples

//'Globalna' zmienna
var owca = "Owca"
var UserName string

//Owca jest wyeksportowaną zmienną, będzie dostępna dla innych paczek
var Owca = "Baran"

//Deklarowanie wielu zmiennych naraz
var (
	punkty = 1
	tygrys = "Marek"
	user   = "Matrix"
	prawda = false
)

var exampleVar string = "Pietrek"

//Zmienna health jest typu int64 i nie ma przypisanej wartośći
var health int64

//Stałe w go mogą być typu: string, bool, numeric(dowolny typ liczbowy jak float)
//Stała
const stalaOwca = "Mirek"

//Baran jest wyeksportowaną stała
const Baran = "Małpa"

//Tworzenie wielu stałych naraz
const (
	player = "Player"
	Gracz         = "Gracz"
)



