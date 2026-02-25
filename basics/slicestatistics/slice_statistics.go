package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	//var array[5] int
	//var array =[...]int{1,2,3,4,5}
	var array=make([]int,0)
	reader:=bufio.NewReader(os.Stdin)
	count,_:=reader.ReadString('\n')
	intcount,_:=strconv.ParseInt(strings.TrimSpace(count),10,64)
	for i:=0; i<int(intcount); i++{
		x,_:=reader.ReadString('\n')
		xint,_:=strconv.ParseInt(strings.TrimSpace(x),10,64)
		array = append(array, int(xint))
	}

	var sum int
	var max int =array[0]
	var min int=array[0]

	for i:= range array{
		sum+=array[i]
		if array[i]>max{
			max=array[i]
		}
		if array[i]<min{
			min=array[i]
		}
		fmt.Print(array[i]," ")
	}

	avg:=float64(sum)/float64(len(array))

	fmt.Print("\nThe sum of these numbers is ",sum)
	fmt.Print("\nThe min of these numbers is ",min)
	fmt.Print("\nThe max of these numbers is ",max)
	fmt.Printf("\nThe avg of these numbers is %.2f",avg)
}