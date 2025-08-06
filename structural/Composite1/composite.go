
package main

import "fmt"



type Employee interface {
	ShowDetails(indent string)
}


type Developer struct {
	Name     string
	Position string
}

func (d *Developer) ShowDetails(indent string) {
	fmt.Printf("%s- %s (Developer)\n", indent, d.Name)
}



type Designer struct {
	Name     string
	Position string
}

func (d *Designer) ShowDetails(indent string) {
	fmt.Printf("%s- %s (Designer)\n", indent, d.Name)
}


type Manager struct {
	Name        string
	Position    string
	Subordinates []Employee
}

func (m *Manager) AddSubordinate(e Employee) {
	m.Subordinates = append(m.Subordinates, e)
}

func (m *Manager) ShowDetails(indent string) {
	fmt.Printf("%s+ %s (Manager)\n", indent, m.Name)
	for _, e := range m.Subordinates {
		e.ShowDetails(indent + "  ")
	}
}


func main() {
	
	dev1 := &Developer{Name: "Sibi"}
	dev2 := &Developer{Name: "Tharshan"}
	designer := &Designer{Name: "MAxi"}

	
	teamLead := &Manager{Name: "kalai"}
	teamLead.AddSubordinate(dev1)
	teamLead.AddSubordinate(dev2)

	
	deptHead := &Manager{Name: "Kumar"}
	deptHead.AddSubordinate(teamLead)
	deptHead.AddSubordinate(designer)

	deptHead.ShowDetails("")
}
