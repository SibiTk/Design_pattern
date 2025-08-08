package main

import "fmt"

type PC struct{
	Model string
	RAM  int
	SSD int
	GPU  string
}

type PCBuilder struct{
	Model string
	RAM  int
	SSD int
	GPU  string
}

func NewPCBuilder(model string)*PCBuilder{
	return &PCBuilder{
		Model: model,
	}
}

// func(p *PCBuilder)SetRAM(ram int)*PCBuilder{
// 	return &PCBuilder{
// 		RAM: ram,
// 	}
// }
// func(p *PCBuilder)NewPCBuilder(model string)*PCBuilder{
// 	p.Model=model
// 	return p
// }


func(p *PCBuilder)SetRAM(ram int)*PCBuilder{
	p.SSD=ram
	return p
}
func(p *PCBuilder)SetSSD(ssd int)*PCBuilder{
	p.SSD=ssd
	return p
}

func(p *PCBuilder)AddGPU(gpu string)*PCBuilder{
	p.GPU=gpu
	return p
}

func(p *PCBuilder)Build()*PC{
	return &PC{
		Model: p.Model,
		RAM: p.RAM,
		SSD: p.SSD,
		GPU: p.GPU,

	}
}

func (p *PC)Print(){
	fmt.Println(("<====PC Details=====>"))
	fmt.Printf("Model : %s\n",p.Model)
	fmt.Printf("RAM :%d\n",p.RAM)
	fmt.Printf("SSD : %d\n",p.SSD)
	fmt.Printf("GPU :%s\n",p.GPU)

}

func main(){
	pc1 := NewPCBuilder("Intel i9").
    SetRAM(32).
    SetSSD(1).
    AddGPU("NVIDIA RTX 4080").
    Build()
	pc1.Print()
}
