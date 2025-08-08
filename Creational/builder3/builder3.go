package main

import "fmt"


type User struct {
	Name    string
	Email   string
	Age     int
	Address string
	Phone   string
}


type UserBuilder struct {
	Name    string
	email   string
	age     int
	address string
	phone   string
}


func NewUserBuilder(name, email string) *UserBuilder {
	return &UserBuilder{
		Name:  name,
		email: email,
	}
}


func (b *UserBuilder) SetAge(age int) *UserBuilder {
	b.age = age
	return b
}

func (b *UserBuilder) SetAddress(address string) *UserBuilder {
	b.address = address
	return b
}

func (b *UserBuilder) SetPhone(phone string) *UserBuilder {
	b.phone = phone
	return b
}

func (b *UserBuilder) Build() *User {
	return &User{
		Name:    b.Name,
		Email:   b.email,
		Age:     b.age,
		Address: b.address,
		Phone:   b.phone,
	}
}


func (u *User) Print() {
	fmt.Println("User Profile:")
	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Email: %s\n", u.Email)
	if u.Age > 0 {
		fmt.Printf("Age: %d\n", u.Age)
	}
	if u.Address != "" {
		fmt.Printf("Address: %s\n", u.Address)
	}
	if u.Phone != "" {
		fmt.Printf("Phone: %s\n", u.Phone)
	}
	fmt.Println()
}

func main() {

	user1 := NewUserBuilder("sibi", "tksibi@gmail.com").Build()


	user2 := NewUserBuilder("Dharmesh", "dharmesh@gmail.com").
		SetAge(30).
		SetAddress("123  Street").
		SetPhone("123-456-78910").
		Build()


	user1.Print()
	user2.Print()
}

