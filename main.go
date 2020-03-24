package main

import (
	"fmt"
	"skeleton/db"
)

func main() {
	fmt.Println("Hello World")
	db.Connect("localhost", "root", "123456", "test", "mysql", "3306")
}
