package main

//In go, you can attach functions to your custom types. They are then referred to as methods.
//In kava, method is a member of a call, in go, method is a member of a type.
//A struct is a custom type.

import (
	"fmt"
)

func main(){
	poodle:=Dog{"Poodle","Bark"}
	poodle.Speak()
	poodle.Sound="ouwwww"
	poodle.Speak()

	println(poodle.SpeakThreeTimes())
}

type Dog struct{
	Breed string
	Sound string
}
//to add a method to a type, you have to declare a new function.
//we are trying to attach types(structs) to a method. for that before the name of the function, add a () and declare a variable and the type you are attaching the method to.
func (d Dog) Speak(){
	fmt.Printf("The %v says %v \n",d.Breed,d.Sound)
}

//Sprintf does the same thing as printf, it outputs a string with placeholders, but instead of outputting to the console, it returns a variable.
func (d Dog) SpeakThreeTimes() string{
	output:=fmt.Sprintf("%v! %v! %v!",d.Sound,d.Sound,d.Sound)
	return output
}

//Methods are essentially functions, but you attach them to your custom types as the receiver of the function. and then within the function, you can use everything else thats available from the object.