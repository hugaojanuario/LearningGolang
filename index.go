package main

import (
	"fmt"
)

type User string

var (
	name User 
	age User
	profission User
)




func main(){
	
	name = "Hugo"
	age = "20"
	profission = "dev backend"


	fmt.Println(name,age,profission)
}
