package main

import (
	"fmt"
	"maps"
	"math"
	"time"
	"unicode/utf8"
)

func main() {
	//forLoop()
	//ifElse()
	//switchingWithInterfaceBonus()
	//arrayKoPo()
	//slicesOfApples()
	//mapsOfBohol()
	//variadicFunc()
	//closures()
	//recursion()
	//rangeRoverCharot_overBuiltinTypes()
	//pointers()
	//stringAndRunes()
	//structs()
	//methods()
	//interfaces()
	//enums()
	//structEmbedding()
	//defered()
	generics()
}
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

type List[T any] struct {
	head, tail *element[T]
}

func generics() {
	var s = []string{"foo", "bar", "baz"}
	fmt.Println("Index of baz: ", SlicesIndex(s, "zoo"))

	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	fmt.Println("list: ", lst.AllElements)
}

func defered() {
	a := 10

	defer fmt.Println("defered", a)

	a = 11
	fmt.Println("undefered", a)

	for b := range a {
		fmt.Println(b + 1)
		if b == 3 {
			defer slicesOfApples() //let this function finishes before going back to the loop
		}
	}

}

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("%d", b.num)
}

type container struct {
	base
	str string
}

func structEmbedding() {
	co := container{
		base: base{
			num: 1,
		},
		str: "str",
	}

	fmt.Printf("co={num:%v}, str: %v\n", co.num, co.str)
	fmt.Println("also num: ", co.base.num)
	fmt.Println("describe: ", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Printf("co={num:%v}, str: %v\n", d.describe(), d.describe())
}

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "Idle",
	StateConnected: "Connected",
	StateError:     "Error",
	StateRetrying:  "Retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Sprintf("unknown server state: %v", s))
	}
}
func enums() {
	ns := transition(StateError)
	fmt.Println("idle state:", ns)

	ns2 := transition(ns)
	fmt.Println("idle state:", ns2)

	ns3 := transition(ns2)
	fmt.Println("idle state:", ns3)
}

type shape interface {
	area() float64
	perimeter() float64
}
type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}
func (c circle) perimeter() float64 {
	return 2 * (c.radius * c.radius)
}
func printing(s shape) {
	if r, ok := s.(rectangle); ok {
		fmt.Println("Rectangle:", r.width, r.height)
	} else if c, ok := s.(circle); ok {
		fmt.Println("Circle:", c.radius)
	}
	fmt.Print(s, " ")
	fmt.Print(s.area(), " ")
	fmt.Println(s.perimeter(), "\n")
}
func detectCircle(s shape) {
	if c, ok := s.(circle); ok {
		fmt.Println("Circle:", c.radius)
	}
}
func detectRect(s shape) {
	if r, ok := s.(rectangle); ok {
		fmt.Println("Rectangle:", r.width, r.height)
	}
}
func interfaces() {
	r := rectangle{width: 10, height: 5}
	c := circle{radius: 5}

	printing(r)
	printing(c)

	detectCircle(r)
	detectCircle(c)
	detectRect(r)
	detectRect(c)

}

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}
func (r rect) perimeter() int {
	return 2*r.width + 2*r.height
}
func methods() {
	r := rect{width: 10, height: 5}
	fmt.Println(r.area(), r.perimeter())
	fmt.Println("Area: ", r.area())
	fmt.Println("Perimeter: ", r.perimeter())

	rp := &r
	rp.width = 20
	rp.height = 30
	fmt.Println("area: ", rp.area())
	fmt.Println("Perimeter: ", rp.perimeter())
}

type Person struct {
	name string
	age  uint8
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 42
	return &p
}
func structs() {
	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{"Alice", 20})
	fmt.Println(&Person{"Fb", 34})
	fmt.Println(newPerson("Jon"))

	s := Person{"JOSE", 28}
	fmt.Println(s.name)

	s1 := &s
	fmt.Println(s1.age)

	s1.age = 51
	fmt.Println(s1.age)

	dog := struct {
		name    string
		goodBoy bool
		dogYear uint8
	}{
		"MAXIE",
		true,
		20}
	fmt.Println(dog)
}
func stringAndRunes() {
	const s = "สวัสดี"
	fmt.Println(s)

	for a := range len(s) {
		fmt.Printf(" %x ", s[a])
	}
	fmt.Println()
	fmt.Println("Rune Count: ", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U at %d:%d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}

}
func examineRune(r rune) {
	if r == 't' {
		fmt.Printf("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

func pointers() {
	i := 1
	fmt.Println(i)

	var zeroVal func(n int) int
	zeroVal = func(n int) int {
		n = 0
		return n
	}
	var zeroPtr func(n *int) int
	zeroPtr = func(i *int) int {
		*i = 0
		return (*i)
	}
	fmt.Println(zeroVal(i))
	fmt.Println(zeroPtr(&i))
	fmt.Println(zeroPtr(&i))
}
func rangeRoverCharot_overBuiltinTypes() {
	// range iterates over elements in a variety of built-in data structures.
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	sum := 0
	for i, num := range nums {
		sum += num
		fmt.Println("Sum:", sum)
		fmt.Println("Num:", i)
	}
}

// recursion start
func fact(n int) int {
	if n == 0 {
		return 1
	}
	//fmt.Println(n)
	//fmt.Printf("fact: %d\n)", uint8(n*fact(n-1)))
	return n * fact(n-1)
}
func recursion() {
	fmt.Println(fact(5))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}

//closure start - o supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func closures() {
	//This function value captures its own i value,
	//which will be updated each time we call nextInt.
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

// closure end
// variadic start
func sum(numbers ...int) {
	fmt.Print(numbers, " = ")
	total := 0
	for _, number := range numbers {
		total += number
	}
	fmt.Println(total)
}
func variadicFunc() {
	sum(1, 2)
	sum(1, 2, 3)
	sum(1, 2, 3, 4)
	sum()
	sum(5, 5, 5, 5)
}

// variadic end
func mapsOfBohol() {
	//To create an empty map, use the builtin
	//make: make(map[key-type]val-type).
	map1 := make(map[string]float32)
	map1["a"] = 1.3243
	map1["b"] = 2.432432
	fmt.Println("map1: ", map1)
	v1 := map1["c"]
	fmt.Println("v1: ", v1)
	fmt.Println("map1 Lenght: ", len(map1))
	delete(map1, "a")
	fmt.Println("map1 Lenght: ", len(map1))
	clear(map1)
	fmt.Println("map1 Lenght: ", len(map1))
	_, prs := map1[`a`]
	fmt.Println("map1[\"a\"]: ", prs)
	//declaring and initializing a new map in the same line
	n := map[string]float32{"Pi": 3.1414, "eulers": 1.23232}
	fmt.Println("map1[]: ", n)

	n2 := map[string]float32{"Pi": 3.1414, "eulers": 1.23232}
	if maps.Equal(n, n2) {
		fmt.Println("maps equal")
	}

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
