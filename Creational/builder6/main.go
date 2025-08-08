package main

import "fmt"

type User struct{
	Name string
	Email string
	Age int
	Status bool 
}

type UserBuilder struct{
	name string
	email string
	age int
	status bool 
}

func  NewUserBuilder(Name,Email string)*UserBuilder{
	return &UserBuilder{
		name:Name,
		email:Email,
	}
}



func (u *UserBuilder) SetAge(age int)*UserBuilder{
u.age =age
return u
}

func (u *UserBuilder)SetTestFlag(Status bool)*UserBuilder{
	u.status=Status
	return u
}
func (u *UserBuilder)Build()*User{
	return &User{
		Name: u.name,
		Email: u.email,
		Age: u.age,
		Status: u.status,

	}
}


func(U *User) Print(){

	fmt.Println("<======testing with Dummy Data======> ")
	fmt.Printf("User : %s\n",U.Name)
	fmt.Printf("Email : %s\n",U.Email)
	fmt.Printf("age : %d\n",U.Age)
	fmt.Printf("Status : %t\n",U.Status)

}

func main(){
	testUser := NewUserBuilder("TestUser", "test@gmail.com").
    SetAge(99).
    SetTestFlag(true).
    Build()
	testUser.Print()

}