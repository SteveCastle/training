package main

import (
	"fmt"
)
// Go does not have classes. Structs are used to create
// compound data structures. Casing is used to hide or export fields.
type car struct{
	Seats int
	Wheels int
	Horn int
}

func main(){
	fmt.Println("Hello world")
}