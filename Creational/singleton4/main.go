package main

import "fmt"

type singleton struct{
connection string
}
var instance *singleton
func GetInstance() *singleton{
	if instance ==nil{
		fmt.Println("Creating single Instance Now....")
		instance = &singleton{connection: "DB Connection"}
	}
	return instance
}


func main(){
db1 :=GetInstance()
fmt.Println("First Call :",db1.connection)
db2 :=GetInstance()
fmt.Println("Second Call : ",db2.connection)
if db1 == db2 {
		fmt.Println("Both instances are the same (singleton verified).")
	} else {
		fmt.Println("Different instances (singleton failed).")
	}

}
