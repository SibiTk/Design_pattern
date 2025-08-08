package main
import "fmt"
type Onlinetest interface{
	DisplayQuestion()
	SelectAnswer()
	Submit()
	Result()
}

func Assessment(o Onlinetest){
		o.DisplayQuestion()
		o.SelectAnswer()
		o.Submit()
		o.Result()
}

type Base struct{}

func (b *Base)Submit(){
	fmt.Println("Answer is Submitted")
}

func (b *Base)Result(){
	fmt.Println("Result is Displayed in the Computer/Laptop")
}

type MCQ struct{
	Base
}

func(m *MCQ)DisplayQuestion(){
	fmt.Println("Displaying MCQ Questions")
}

func (m *MCQ)SelectAnswer(){
	fmt.Println("Choose the Best Answer")
}

type TrueOrFalse  struct{
	Base
}
func(t *TrueOrFalse)DisplayQuestion(){
	fmt.Println("Displaying True or False Questions")
}

func (t *TrueOrFalse)SelectAnswer(){
	fmt.Println("Select True Or False")
}

func main(){
	fmt.Println("<=====MCQ Assesment======>")
	Assessment(&MCQ{})
	fmt.Println("<======True Or False======>")
	Assessment(&TrueOrFalse{})
}