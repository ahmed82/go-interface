package main

import "fmt"

type Tuts interface {
}

func arrayd() {
	x := []int{0, 1, 2, 3, 4, 5}
	for index, value := range x {
		fmt.Println("Index: ", index, " Value: ", value)
	}
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func pointer() {
	a := [...]int{1, 2, 3} //the {...} create array size enouph to hold the data that initialais
	b := &a
	b[1] = 5
	fmt.Println("############################")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Length: %v\n", cap(a))
}

func slic() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //tur the slice to arrays
	b := a[:]   //slice of all elements of a
	c := a[3:]  //slice from 4th element to the end
	d := a[:6]  //slice first 6 elements
	e := a[3:6] // slice the 4th, 5th and 6th elements
	a[5] = 42
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
