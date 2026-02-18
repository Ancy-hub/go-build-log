package main

import (
	"fmt"
	"math"
	"time"
)

//in go, you cant add int to a float
//wrap value in target types as a function call
//to declare multiple: i1,i2,i3=2,3,4
//round fucntion rounds the value to nearest integer.
func main() {
	var x int = 5
	var y float64 = 4
	sum := float64(x) + y
	fmt.Printf("The sum is %v, Type: %T\n", sum, sum)

	sum=math.Round(sum*100)/100
	fmt.Print(sum,"\n")

	//for date and time
	t:=time.Date(2002,time.March,26,13,0,0,0,time.UTC)
	fmt.Print(t,"\n")

	n:=time.Now()
	fmt.Print(n)
	//n.Format(time.ANSIC)
	//n.AddDate(0,0,1)->prints tmrw.
	
}