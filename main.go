package main

import "fmt"

type Printer interface {
	Print() string
}
type myPringter struct{}

func (mp myPringter) Print() string {
	return "Printing One page ..."
}

type scondPrinter struct{}

func (sp scondPrinter) Print() string {
	return "Printing More pages >>>"
}

//func process(equipment myPringter) {
func process(equipment Printer) {
	fmt.Println("Running Print,,,", equipment.Print())
}
func main() {
	printer := myPringter{}
	otherPrinter := scondPrinter{}
	process(printer)
	process(otherPrinter)
}
