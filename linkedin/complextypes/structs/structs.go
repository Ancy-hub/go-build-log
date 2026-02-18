package main

//Struct are similar to functions. they encapsuate both data and methods. but go doesnt have inheritance model.
//each struct is independent with its own fields for data management.
//struct name has uppercase character, which makes it public to the rest of the application.
import "fmt"

func main(){
	fmt.Println("Structs in go")

	poodle:=Dog{"Poodle",34}
	fmt.Println(poodle)
	fmt.Printf("%+v\n",poodle)//a %+v menas include both name and values of the fields. 
	fmt.Printf("Breed: %v\n",poodle.Breed)
	poodle.Weight=29
	fmt.Println(poodle)
}

type Dog struct{
	Breed string
	Weight int

}