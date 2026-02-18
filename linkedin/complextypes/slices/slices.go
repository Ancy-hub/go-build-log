package main

//A slice in go is an abstraction layer that sits on the top of an array.
//When you declare a slice, the runtime allocates the required memory and creates the underlying array of data in the background and returns a reference to the requested slice.
//all value will be of same type, but can be sorted and added and removed.
// simplest syntax is just to declare like array but put in a pair of brackets without a number.

import (
	"fmt" 
	"sort"
)

func main(){
	//basic declaration
	//var  colors=[]string{"Red","Pink","Green"}
	//Memory efficient way of declaring slice is
	var colors= make([]string,0,3)//the make function decalres the iteam and allocates memory in the background at the same time
	//the 0 is the initial size and the 3 is the capacity of the slice, if the capacity is reached, the slice will double the size of the array.
	fmt.Println(colors)
	colors=append(colors,"Red","Green","Pikn","aNCCC")
	fmt.Println(len(colors),cap(colors))
	fmt.Println(remove(colors,0))
	//sorting
	sort.Strings(colors)
}

//removing from slice is bit comples, and best done using an additional function
func remove(slice []string, i int) []string{
	return append(slice[:i],slice[i+1:]...)//this is the syntax for removing an element from a slice])

}
