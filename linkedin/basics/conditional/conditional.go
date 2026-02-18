package main

import (
	"fmt"
	"time"
)

//no need for paranthesis. go main goal is to reduce typing.
func main(){
	// //answer:=43;
	// var result string
	// if answer:=42; answer<0{
	// 	result="answer less than 0"
	// }//else is optional
	// fmt.Println(result)
	//switch
	weekday:=time.Now().Weekday()
	fmt.Println(weekday)
	dayNumber:=int(weekday)
	var result string
	switch dayNumber{
	case 1:
		result = "Monday"
	}
	fmt.Println(result)

	x:=-42
	switch{
	case x<0:
		result = "less than 0"
	case x>=0:
		result = "greated than 0"
	default:
		result="zero"
	}
	fmt.Println(result)
	//for
	value:=0
	sum:=0
	for value<4{
		sum+=value
		value++
	}
	fmt.Println(sum)
	//break type
	sum =1
	for sum<1000{
		sum+=sum
		if sum>200{
			goto theEnd
		}
	}

	theEnd: println("end")


}