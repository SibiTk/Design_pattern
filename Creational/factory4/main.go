package main

import "fmt"

type  DB interface {
	Connect() string
}

type MySQL struct{}

func (m MySQL) Connect() string {
    return "Connected toMySQL database"
}

type Postgres struct{}

func (p Postgres) Connect() string {
    return "Connected to Postgres Database"
}

type SQLite struct{}

func (s SQLite) Connect() string {
    return "Connected to sqlite database"
}


func DBFactoctory(dbType string)DB{
	switch dbType {
	case "mysql":
		return MySQL{}
	case "Postgres":
		return Postgres{}
	case "SQLite":
		return SQLite{}
	}
	return nil
}

func main(){
   postgres :=DBFactoctory("Postgres")
   fmt.Println(postgres.Connect())
}