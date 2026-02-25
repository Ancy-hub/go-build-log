package main

//read a text file, count number of words,lines,characters, print results. hint: os.open, bufio.scanner,file read loops
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	file,err:=os.Open("./sample.txt")
	if err!=nil{
		fmt.Println("Error:",err)
		return
	}
	defer file.Close()
	scanner:=bufio.NewScanner(file)
	var chars,words,lines int
	for scanner.Scan(){
		line:=scanner.Text()
		lines++
		words+=len(strings.Fields(line))
		chars+=len(line)+1
	}
	fmt.Printf("Lines: %d\nWords: %d\nCharacters: %d\n", lines, words, chars-1)
}