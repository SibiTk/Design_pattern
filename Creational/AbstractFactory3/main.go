package main

import "fmt"

type  GUIFactory interface{
	CreateButton() Button
	CreateCheckbox() Checkbox
}
type Button interface{
	press() string
}
type Checkbox interface{
	Check() string
}
//concrete product for windows
type WindowsButton struct {}
type WindowsCheckbox struct {}
//concrete product  for mac
type MacButton struct{}
type MacCheckbox struct {}
//factories
type WindowFactory struct{}
type MacFactory struct{}

func (w *WindowsButton)press() string{return "Windows Button is pressed"}
func (w *WindowsCheckbox)Check() string{return "Windows CheckBox Clicked"}

func (m *MacButton)press()string {return "Mac button is pressed"}
func (m  *MacCheckbox)Check()string {return "Mac CheckBox is Clicked"}

func (w *WindowFactory) CreateButton() Button   { return &WindowsButton{} }
func( w *WindowFactory)CreateCheckbox() Checkbox {return &WindowsCheckbox{}}

func (m *MacFactory)CreateButton() Button {return &MacButton{}}
func (m *MacFactory)CreateCheckbox() Checkbox { return &MacCheckbox{}}


func main(){
	 var wf GUIFactory = &WindowFactory{}
	 button := wf.CreateButton()
 checkbox := wf.CreateCheckbox()
 fmt.Println(button.press())
 fmt.Println(checkbox.Check())
 var mf GUIFactory =&MacFactory{}
    button =mf.CreateButton()
	checkbox =mf.CreateCheckbox()
	fmt.Println(button.press())
	fmt.Println(checkbox.Check())
	
}
