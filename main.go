package main

import (
	"fmt"

	"github.com/ahmed82/go-interface/pkg"
	service "github.com/ahmed82/go-interface/pkg/services"
	//"github.com/ahmed82/go-interface/pkg/service"
)

type Printer interface {
	Print() string
}

type Scanner interface {
	Scan() string
}

type Faxer interface {
	Fax() string
}

//Composit interface that combain TWO interfaces
type PrinterScanner interface {
	Printer
	Scanner
}

type FaxerPrinterScanner interface {
	Faxer
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

func (mp myPrinter) Fax() string {
	return "Faxed One page ..."
}

type scondPrinter struct{}

func (sp scondPrinter) Print() string {
	return "Printing More pages >>>"
}

func (sp scondPrinter) Scan() string {
	return "Scanned More pages >>>"
}

func (sp scondPrinter) Fax() string {
	return "Faxed More pages >>>"
}

//func process(equipment myPringter) {
//func process(equipment Printer) {
func process(equipment FaxerPrinterScanner) {
	fmt.Println("Running Print,,,", equipment.Print())
	fmt.Println("Running Scan,,,", equipment.Scan())
	fmt.Println("Sending Fax,,,", equipment.Fax())
}
func main() {
	printer := myPrinter{}
	otherPrinter := scondPrinter{}
	process(printer)
	process(otherPrinter)
	tutorials()
}

func tutorials() {
	//slic()
	//pointer()
	fmt.Println("---------------------------------------")
	//arrayd()
	fmt.Println("---------------------------------------")
	//composit()
	fmt.Println("---------------------------------------")
	//getTag()
	fmt.Println("-------------************---------------")
	//errHandel()
	//goRoutin()
	// channel()
	//channel2()

	//pkg.Robot()

	//pkg.PanicEx()
	pkg.TestDivide()
	service.InterfaceAdv()
}
