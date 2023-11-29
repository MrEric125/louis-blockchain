package main

import "fmt"

type Mac struct {

	printer Printer

}

func (m *Mac) Print()  {
	fmt.Println("print request for mac")

	m.printer.PrintFile()

}

func (m *Mac) SetPrinter(p Printer)  {

}
