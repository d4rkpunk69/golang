package main

import (
	"fmt"
	"time"
)

func main() {
	//forLoop()
	//ifElse()
	//switchingWithInterfaceBonus()
	//arrayKoPo()
	//slicesOfApples()
	mapsOfBohol()
}

func mapsOfBohol() {
	//To create an empty map, use the builtin
	//make: make(map[key-type]val-type).
	map1 := make(map[string]int)
	map1["a"] = 1
	map1["b"] = 2
	fmt.Println("map1: ", map1)
}
func slicesOfApples() {
	var slice1 []string
	fmt.Printf("Uninitializes slice1 %v : %s : %t\n", len(slice1), slice1, len(slice1) == 0)

	slice2 := make([]string, 3)
	fmt.Printf("empty: %s, len: %d, cap: %d\n", slice2, len(slice2), cap(slice2))
	slice2[0] = "Hello"
	slice2[1] = "F'n"
	slice2[2] = "World"
	for a := range len(slice2) {
		fmt.Printf("slice2[%d]: %s\n", a, slice2[a])
	}
	fmt.Printf("len: %d, cap: %d\n", len(slice2), cap(slice2))
	slice2 = append(slice2, "This")

	fmt.Printf("+ len: %d, + cap: %d\n", len(slice2), cap(slice2))
	slice2 = append(slice2, "Is your")
	fmt.Printf("%s worst enemy\n", slice2[4])
	fmt.Printf("++ len: %d, ++ cap: %d\n", len(slice2), cap(slice2))

	slice3 := make([]string, len(slice2))
	copy(slice3, slice2)
	fmt.Printf("copy: %s\n", slice3)
	fmt.Printf("slice3 len: %d, cap: %d\n", len(slice3), cap(slice3))

	slicing := slice3[2:4]
	fmt.Printf("slicing %s \n", slicing)

	slice4 := []string{"a", "b", "c", "d", "This"}
	fmt.Printf("slice4 len: %d, cap: %d\nvalues: %s\n", len(slice4), cap(slice4), slice4)
	for b := range len(slicing) {
		for c := range len(slice4) {
			fmt.Printf("b: %d, c: %d, slicing: %s, slice4 %s\n", b, c, slicing[b], slice4[c])
			if slicing[b] == slice4[c] {
				fmt.Printf("Found it %s : %s\n", slicing[b], slice4[c])
			}
		}
	}
}
func arrayKoPo() {
	var myArray [5]int
	fmt.Printf("myArray length: %v\n", len(myArray))
	myArray[4] = 10000
	fmt.Printf("myArray val last: %v\n", myArray[4])
	myArray2 := [...]string{"James", "Bond", "Shaken, not stirred"}
	fmt.Printf("myArray2 length: %v\n", len(myArray2))
	fmt.Printf("myArray2's val: %s\n", myArray2)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		fmt.Printf("i:%v\n", i)
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Printf("twoD length: %v\n", len(twoD))
	fmt.Printf("twoD val: %v\n", twoD)
	twoD = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("twoD2 length: %v\n", twoD)
}
func switchingWithInterfaceBonus() {

	var helYeah string = "fuckin hellf"
	switch len(helYeah) {
	case 10:
		fmt.Println("NOOOOO")
	case 11:
		fmt.Println(helYeah)
	default:
		fmt.Println("WA KA TAGNAA")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	case t.Hour() < 19:
		fmt.Println("Good evening")
	default:
		fmt.Println("Aslept")
	}

	//I am now understading an interface
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		default:
			fmt.Println("I'm nothing", t)
		}
	}
	whatAmI(1)
	whatAmI("hey")
	whatAmI(3.22)
	whatAmI(0)
}
func ifElse() {
	var helNo string = "hell noo0000"
	if a := len(helNo); a < 10 {
		fmt.Println(a)
	}
}
func forLoop() {

	var myString string = "Hello World"
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	for j := 0; j < i; j++ {
		fmt.Println(j)
	}
	for i := range len(myString) {
		fmt.Printf("%c", myString[i])
		fmt.Println("range", i)
		if myString[i] == 'W' {
			break
		}
	}
	for n := range len(myString) {
		if n%2 == 0 {
			fmt.Printf("Skipped: %c \n", myString[n])
			continue
		}
		fmt.Printf("Unskipped: %c \n", myString[n])
	}

}
