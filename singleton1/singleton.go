
package main
import "fmt"
type Singleton struct{
    name string
}
var(
     instance *Singleton
     
     
)

func GetInstance() *Singleton  {
    if instance == nil {
        instance = &Singleton{name:"My name is Sibi"}
    }
    return instance
}
func main() {
    a := GetInstance()
    b := GetInstance()
 
    fmt.Println(a == b)
    fmt.Println(a)
    fmt.Println(b)
}
