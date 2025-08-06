package main




import (
 "fmt"
)

type OldSocket struct{}

 type Adapter struct {
  OldSocket *OldSocket
 }


  type TypeCPlug interface {
  Insert() string
 }


func (o *OldSocket) PlugIn() string {
 return "Power from old socket"
}

func (a *Adapter) Insert() string {

 return a.OldSocket.PlugIn() + " adapted to Type-C"
}

func main() {
 s := &OldSocket{}
 adapter := &Adapter{OldSocket: s}

 fmt.Println(adapter.Insert()) 
}
