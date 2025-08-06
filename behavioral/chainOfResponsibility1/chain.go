package main

import "fmt"


type Department interface{
	Execute(*Patient)
	SetNext(Department)
}
 
type BaseDepartment struct{
	next Department
}

type Patient struct{
	name string
	checkIn bool
	diagnosed bool
	tested bool
	PaymentCompleted bool
}

func (b *BaseDepartment) SetNext(next Department){
	b.next =next
}
func (b * BaseDepartment)  CallNext(p *Patient){
	if b.next !=nil{
		b.next.Execute(p)
	}

}
type Reception struct {
	BaseDepartment
}

func ( r *Reception)Execute(p *Patient){
	if p.checkIn{
		fmt.Printf("Reception : Patient %s already Checked in\n",p.name)
		r.CallNext(p)
		return
	}
	fmt.Printf("Reception :Checking In Patient name : %s\n",p.name)
	p.checkIn= true
	r.CallNext(p)
}
type Doctor struct {
	BaseDepartment
}

func (d *Doctor) Execute(p *Patient){
	if p.diagnosed{
		fmt.Printf("Doctor : Patient %s  Already Diagnosed.\n",p.name)
		d.CallNext(p)
		return
	}
	fmt.Printf("Doctor : Diagnosing Patient Name : %s \n",p.name)
	p.diagnosed= true
	d.CallNext(p)
}

type Medical struct{
	BaseDepartment

}
func (m *Medical)Execute(p *Patient){
	if p.tested{
		fmt.Printf("Medical Exam : Patiernt %s  Already Tested \n",p.name)
		m.CallNext(p)
		return
	}
	fmt.Printf("MEdical Exam : Performing tests for the patient %s..... \n",p.name)
	p.tested=true
	m.CallNext(p)
}
type Cashier struct{
	BaseDepartment
}
func (c *Cashier) Execute(p *Patient){
	if p.PaymentCompleted{
		fmt.Printf("Payment is already Completed for %s\n",p.name)
		c.CallNext(p)
		return
	}
	fmt.Printf("CAshier : Processing Payment for %s...\n",p.name)
	p.PaymentCompleted=true
}

func main(){
	reception :=&Reception{}
	doctor :=&Doctor{}
	medical :=&Medical{}
	cashier :=&Cashier{}
	reception.SetNext(doctor)
	doctor.SetNext(medical)
	medical.SetNext(cashier)


	patient1 :=&Patient{name : "Sibi T"}
	patient2 := &Patient{name:"Maxi"}
	reception.Execute(patient1)
	reception.Execute(patient1)
	reception.Execute(patient2)


}