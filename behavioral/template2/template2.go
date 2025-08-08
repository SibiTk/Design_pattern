package main

import (
	"fmt"
)

type ReportGenerator interface {
	Prepare()
	GenerateContent()
	Finish()
}


func GenerateReport(r ReportGenerator) {
	r.Prepare()
	r.GenerateContent()
	r.Finish()
}


type PDFReport struct{}

func (p PDFReport) Prepare() {
	fmt.Println("Preparing PDF Report...")
}

func (p PDFReport) GenerateContent() {
	fmt.Println("Generating PDF content...")
}

func (p PDFReport) Finish() {
	fmt.Println("Finalizing PDF Report.")
}


type ExcelReport struct{}

func (e ExcelReport) Prepare() {
	fmt.Println("Preparing Excel Report...")
}

func (e ExcelReport) GenerateContent() {
	fmt.Println("Generating Excel content...")
}

func (e ExcelReport) Finish() {
	fmt.Println("Finalizing Excel Report.")
}


type HTMLReport struct{}

func (h HTMLReport) Prepare() {
	fmt.Println("Preparing HTML Report...")
}

func (h HTMLReport) GenerateContent() {
	fmt.Println("Generating HTML content...")
}

func (h HTMLReport) Finish() {
	fmt.Println("Finalizing HTML Report.")
}


func main() {
	fmt.Println("<== PDF Report ==>")
	GenerateReport(PDFReport{})

	fmt.Println("\n<== Excel Report ==>")
	GenerateReport(ExcelReport{})

	fmt.Println("\n<== HTML Report ==>")
	GenerateReport(HTMLReport{})
}
