package main

import "fmt"

type USBPrinter interface {
	PrintFile()
}

type Printer struct{}

func (lp *Printer) PrintLegacyFile() {
	fmt.Println("Printing file using a legacy printer")
}

type PrinterAdapter struct {
	Legacy *Printer
}

func (pa *PrinterAdapter) PrintFile() {

	fmt.Println("Adapter converts USB print to legacy format:")
	pa.Legacy.PrintLegacyFile()
}

func main() {

	Printer := &Printer{}

	adapter := &PrinterAdapter{
		Legacy: Printer,
	}

	printViaUSB(adapter)
}

func printViaUSB(printer USBPrinter) {
	printer.PrintFile()
}
