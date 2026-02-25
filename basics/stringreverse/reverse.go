package main

import (
	"fmt"
)

func main(){
	var s string
	fmt.Print("Type something: ")
	fmt.Scanln(&s)
	for _, char := range s{
		fmt.Println(string(char))
	}

	for i:=len(s)-1;i>=0;i--{
		fmt.Println(string(s[i]))
	}
	
}