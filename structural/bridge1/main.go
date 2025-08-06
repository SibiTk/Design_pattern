package main

import "fmt"


type Printer interface {
	PrintFile()
}


type Epson struct{}
func (p *Epson) PrintFile() {
	fmt.Println("Printing by Epson printer")
}


type HP struct{}
func (p *HP) PrintFile() {
	fmt.Println("Printing by HP printer")
}


type Computer interface {
	Print()
	SetPrinter(Printer)
}

type Mac struct {
	printer Printer
}
func (m *Mac) Print() {
	fmt.Print("Print request for Mac: ")
	m.printer.PrintFile()
}
func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}


type Windows struct {
	printer Printer
}
func (w *Windows) Print() {
	fmt.Print("Print request for Windows: ")
	w.printer.PrintFile()
}
func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

func main() {
	epson := &Epson{}
	hp := &HP{}

	mac := &Mac{}
	mac.SetPrinter(epson)
	mac.Print()     
	win := &Windows{}
	win.SetPrinter(hp)
	win.Print()     
}
