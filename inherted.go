package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func composit() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)
}

type Sea struct {
	Name   string `required max:"100"`
	Origin string
}

func getTag() {
	t := reflect.TypeOf(Sea{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
