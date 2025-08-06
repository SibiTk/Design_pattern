package main

import "fmt"


type Document interface {
	Clone() Document
	Print()
}

type Report struct {
	Title   string
	Content string
}


func (r *Report) Clone() Document {
	clone := *r 
	return &clone
}

func (r *Report) Print() {
	fmt.Printf("Report Title: %s\nContent: %s\n\n", r.Title, r.Content)
}

func main() {
	
	template := &Report{
		Title:   "monthly Sales Report",
		Content: "This is the base content for the monthly sales report.",
	}


	reportJan := template.Clone()
	reportFeb := template.Clone()

	
	rJan := reportJan.(*Report)
	rJan.Title = "Monthly Sales Report - January"
	rJan.Content = "Sales increased by 10% in January."

	rFeb := reportFeb.(*Report)
	rFeb.Title = "Monthly Sales Report - February"
	rFeb.Content = "Sales decreased by 5% in February."


	fmt.Println("<== Base Template ==>")
	template.Print()

	fmt.Println("<== Cloned Reports ==>")
	rJan.Print()
	rFeb.Print()
}
