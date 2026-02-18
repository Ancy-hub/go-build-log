package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	// if you want to ignore an error use _ instead of a name
	str, _:= reader.ReadString('\n')
	fmt.Println(str)

	fmt.Print("Enter a number: ")
	str, _= reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str),64)
	if(err!=nil){
		fmt.Println(err)
	}else{
		fmt.Println("Value of number: ",f)
	}


}