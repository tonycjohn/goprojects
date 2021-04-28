package main

import (
	"fmt"
	"strconv"
)

var i float32 = 3.14

//Doctor Structure
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func playground() {
	answer, err := divideNumbers(25.00, 0.00)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(answer)

	var i int
	i = 42
	var j float32
	j = float32(i)
	k := 99
	var l string
	l = strconv.Itoa(k)
	var n bool
	o := 1 == 1
	p := 1 == 2

	fmt.Println("Hello Go!")
	fmt.Printf("value and type of k: %v,%T \n", k, k)
	fmt.Printf("value and type of i: %v,%T", i, i)
	fmt.Printf("\nvalue and type of j: %v,%T", j, j)
	fmt.Printf("\nvalue and type of l: %v,%T", l, l)
	fmt.Printf("\nvalue and type of o bool: %v - %T, p bool is %v-%T, n bool is %v-%T ", o, o, p, p, n, n)

	//bit operation. works only with integers, not float
	//also % operator works only  with integers
	fmt.Printf("\nbit operation")
	a := 10
	b := 3
	roleByte := 1
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)
	fmt.Println(a << 3) //bit shifting
	fmt.Printf("%b\n", roleByte<<1)
	//float
	c := 3.14e72
	fmt.Println(c)

	//complex numbers
	fmt.Printf("\nComplex Numbers\n")
	var d complex128 = 1 + 2i
	e := 2 + 3i
	fmt.Println(d + e)
	fmt.Println(d - e)
	fmt.Println(d * e)
	fmt.Println(d / e)
	fmt.Printf("%v, %v", real(d), imag(d))

	//string -utf-8 byte
	fmt.Printf("\nbyte and rune")
	s := "hello go!"
	sbyte := []byte(s)
	fmt.Printf("\n%v,%T", s[2], s[2])
	fmt.Printf("\n%v,%T", sbyte, sbyte)

	//rune type alias for int 32. it represents UTF32 type
	var r rune = 'l'
	fmt.Printf("\n%v, %T", r, r)

	//Constants. This is typed constant
	const myConst int = 42 //must be assignable at compile time.
	fmt.Printf("%v, %T", myConst, myConst)

	//untyped constant
	const untypedConst = 42
	fmt.Printf("\n%v, %T", untypedConst, untypedConst)

	//enumerated constant
	//iota is a counter used to create enumerated constant
	const (
		enum1 = iota
		enum2 = iota
		enum3 = iota
	)
	fmt.Printf("\n%v", enum1)
	fmt.Printf("\n%v", enum2)

	const (
		catSpecialist = iota
		dogSpecialist
		chickenSpecialist
	)
	var specialistType int = chickenSpecialist
	fmt.Printf("\n%v", specialistType == chickenSpecialist)

	//Arrays and Slices
	grades := [3]int{33, 24, 48}
	fmt.Printf("\nGrades: %v Length: %v", grades, len(grades))
	newGrades := grades //made a  literal copy
	gradePointer := &grades
	var emptyArray [3]string
	fmt.Printf("\nGrades: %v", newGrades)
	fmt.Printf("\nGrades Pointer: %v", gradePointer)
	fmt.Printf("\nEmpty Array: %v", emptyArray)

	gradeslice := []int{1, 3, 4}
	fmt.Printf("\n Gradeslice %v", gradeslice)
	makeSlice := make([]int, 0, 10)
	makeSlice = append(makeSlice, 1)
	fmt.Printf("\n makeslice value %v, length %v, capacity %v\n", makeSlice, len(makeSlice), cap(makeSlice))

	//Maps and Structs
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12801539,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations)
	//use make function
	makePopulations := make(map[string]int)
	makePopulations = map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12801539,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(makePopulations)
	fmt.Println(makePopulations["Illinois"])

	for k, v := range makePopulations {
		println(k, v)
	}

	//add a new element
	makePopulations["Georgia"] = 10310371
	fmt.Println(makePopulations)
	delete(makePopulations, "Georgia")
	fmt.Println(makePopulations)
	pop, ok := makePopulations["Georgia"] //comma ok syntx to test availability of a key
	fmt.Println(pop, ok)
	fmt.Println(len(makePopulations))
	//map like slice is a referece typs

	numberList() //calling a function

	//struct
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor.actorName)

	//httpCall() //calling http request

	//Interface
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

//INTERFACES

//Writer  interface
type Writer interface {
	Write([]byte) (int, error)
}

//ConsoleWriter impl
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

//Incrementer interface
type Incrementer interface {
	Increment() int
}

//IntCounter impl
type IntCounter int

//Increment impl
func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
