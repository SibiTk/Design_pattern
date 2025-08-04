
package main

import (
	"fmt"
	"time"
)


type Database interface {
	Query(query string) string
}


type RealDatabase struct{}

func (db *RealDatabase) Query(query string) string {
	
	time.Sleep(2 * time.Second)
	return "Result for query: " + query
}


type ProxyDatabase struct {
	realDB   *RealDatabase
	cache    map[string]string
}

func NewProxyDatabase() *ProxyDatabase {
	return &ProxyDatabase{
		cache: make(map[string]string),
	}
}

func (p *ProxyDatabase) Query(query string) string {

	if result, ok := p.cache[query]; ok {
		fmt.Println("Returning cached result for:", query)
		return result
	}

	if p.realDB == nil {
		fmt.Println("Creating RealDatabase instance...")
		p.realDB = &RealDatabase{}
	}

	fmt.Println("Querying real database for:", query)
	result := p.realDB.Query(query)

	p.cache[query] = result
	return result
}

func main() {
	db := NewProxyDatabase()

	fmt.Println(db.Query("SELECT * FROM users"))

	fmt.Println(db.Query("SELECT * FROM orders"))

	
}
