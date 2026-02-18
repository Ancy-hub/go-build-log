package main

//you may face errors while dealing with files in go. so always set up a single way for examining errors.

import (
	"fmt"
	"io"
	"os"
)

func main(){
	fmt.Println("Files")
	fileName:="./fromString.txt"
	file, err := os.Create(fileName)
	checkError(err)
	//always close a file when opened
	defer file.Close() //defer means wait until everything in this function has been executed.
	length, err:=io.WriteString(file,"Hello from go!")
	fmt.Printf("Wrote a file with %v characters\n", length)
	//you will have to always check error from file creation function before closing it.
	readFile(fileName)
}

func checkError(err error){
	if err!=nil{
		panic(err)
	}
}

func readFile(fileName string){
	data,err:=os.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file: ",string(data))
}