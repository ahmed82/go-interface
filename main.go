package main

import "fmt"

type Printer interface {
	Print() string
}

type Scanner interface {
	Scan() string
}

//Composit interface that combain TWO interfaces
type PrinterScanner interface {
	Printer
	Scanner
}

type myPrinter struct{}

func (mp myPrinter) Print() string {
	return "Printing One page ..."
}

func (mp myPrinter) Scan() string {
	return "Scanned One page ..."
}

type scondPrinter struct{}

func (sp scondPrinter) Print() string {
	return "Printing More pages >>>"
}

func (sp scondPrinter) Scan() string {
	return "Scanned More pages >>>"
}

//func process(equipment myPringter) {
//func process(equipment Printer) {
func process(equipment PrinterScanner) {
	fmt.Println("Running Print,,,", equipment.Print())
	fmt.Println("Running Scan,,,", equipment.Scan())
}
func main() {
	printer := myPrinter{}
	otherPrinter := scondPrinter{}
	process(printer)
	process(otherPrinter)
}
