package service

import "fmt"

func InterfaceTester() {
	var w Writer1 = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}

type Writer1 interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
