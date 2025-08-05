package main

import "fmt"

type DataProcessor interface{
	ReadData()
	ProcessData()
	SaveResult()
	Process()

}

type ProcessorTemplate struct{
	processor DataProcessor
}
func (p *ProcessorTemplate) Process() {
	p.processor.ReadData()
	p.processor.ProcessData()
	p.processor.SaveResult()
}

type  CSVProcessor struct{}

func (c *CSVProcessor) ReadData(){
	fmt.Println("Reading the CSV file")
}

func (c *CSVProcessor)ProcessData(){
	fmt.Println("Processing the CSV Data")
}

func(c *CSVProcessor)SaveResult(){
	fmt.Println("Saving CSV Data result")
}

func (c *CSVProcessor)Process(){
	template:=&ProcessorTemplate{processor: c}
	template.Process()
}

type JSONProcessor struct{}

func (j *JSONProcessor) ReadData(){
	fmt.Println("Reading the  JSON file")
}

func (j *JSONProcessor)ProcessData(){
	fmt.Println("Reading the JSON file")
}

func (j *JSONProcessor)SaveResult(){
	fmt.Println("Processing the JSON file")
}

func (j *JSONProcessor)Process(){
	template :=&ProcessorTemplate{processor:j}
	template.Process()
}

func main(){
	fmt.Println("CSV Data Processing")
	csvProc :=&CSVProcessor{}
	csvProc.Process()
	jsonProc :=&JSONProcessor{}
	jsonProc.Process()
}