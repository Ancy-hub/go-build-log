package main

import "fmt"

func main(){
	value:=42
	pointer:=&value
	fmt.Println("Value of pointer is: ",*pointer)

	var p* int=&value
	fmt.Println(p)
	//this prints the memory address of p and value.
}
