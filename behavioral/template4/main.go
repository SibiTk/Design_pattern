package main

import "fmt"


type Recipe interface {
	PrepareIngredients()
	Cook()
	Serve()
}

func PrepareRecipe(r Recipe) {
	r.PrepareIngredients()
	r.Cook()
	r.Serve()
}


type PastaRecipe struct{}

func (p *PastaRecipe)PrepareIngredients(){
	fmt.Println("Prepare the Ingredients for the pasta  to cook")
}
func (p *PastaRecipe)Cook(){
	fmt.Println("Cooking Pasta ")
}
func(p *PastaRecipe)Serve(){
	fmt.Println("Serve the food to the family members")
}

type FriedRiceRecipe struct{}
func (f *FriedRiceRecipe)PrepareIngredients(){
  fmt.Println("Preparing the Fried Rice Ingredients for the cook")
}
func (f *FriedRiceRecipe)Cook(){
	fmt.Println("Cooking Fried Rice")
}
func (f *FriedRiceRecipe)Serve(){
	fmt.Println("Serving the family Members  who need  Fried Rice")
}

func main(){
    fmt.Println("<=== Loves Pasta ===>")
	PrepareRecipe(&PastaRecipe{})
	fmt.Println("<== Love Fried Rice ==>")
	PrepareRecipe(&FriedRiceRecipe{})

}