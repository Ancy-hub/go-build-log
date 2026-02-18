package main
//An array in go has a fixed number of items. Once decalred the size cannot be changed.
//The size of an array is part of its type, so arrays cannot be resized. This seems limiting, Go provides a convenient way of working with arrays.
//An arrays value can be of any type, but cannot mix types.
import "fmt"

func main(){
	var colors[3] string//this is a collection of 3 string values.
	colors[0]="Red"
	colors[1]="Green"
	colors[2]="Pink"
	fmt.Println(colors)

	var numbers=[5] int{1,2,3,4,5}//to declare and initialize an array
	fmt.Println(numbers)

	//In Go, an array is a value. When an array is passed to a function, a copy of the array is passed. 
	//When an array is returned from a function, a copy of the array is returned.
	//But storing data is all you can do with arrays, you cant easily sort, add or remove items at runtime. For those, you should package your order data in something called a slice.
	//To pass an array to a function, pass a pointer to the array.
	//To return an array from a function, return a pointer to the array.

	var a=[...]int{1,2,3,4,5}//this is slice of 5 integers
	fmt.Printf("Type is %T",a)


	//When you use range on a slice, it returns two values:  Index (position in the slice), Value (the actual element)
}