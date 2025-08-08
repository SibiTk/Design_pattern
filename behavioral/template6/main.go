package main

import "fmt"

type BuildHouse interface {
	buildFoundation()
	buildPillars()
	buildWalls()
	buildWindows()
}

func BuildingHouse(b BuildHouse) {
	b.buildFoundation()
	b.buildPillars()
	b.buildWalls()
	b.buildWindows()
}

type BaseFoundation struct{}

func (b *BaseFoundation)buildFoundation() {
	fmt.Println("Building foundation with cement,iron rods and sand")
}

type WoodenHouse struct {
	BaseFoundation
}

func (w *WoodenHouse) buildPillars(){
	fmt.Println("Building Pillars with wooden Coating")
}

func (w *WoodenHouse)buildWalls(){
	fmt.Println("Building  Walls with wooden Coating")
}

func (w *WoodenHouse)buildWindows(){
	fmt.Println("Building Wooden Windows")
}

type GlassHouse struct{
   BaseFoundation	
}

func( g *GlassHouse)buildPillars(){
	fmt.Println("Building Pillars with glass Coating")
}

func (g *GlassHouse)buildWalls(){
	fmt.Println("Building walls with Windows Coating")
}

func (g *GlassHouse)buildWindows(){
	fmt.Println("Building windows with wooden Coating ")
}

func main(){
	fmt.Println("<====Constructing Wooden House====>")
	BuildingHouse(&WoodenHouse{})

	fmt.Println("<=====Constructing Glass House=====>")
	BuildingHouse(&GlassHouse{})
}