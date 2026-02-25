package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Calculator")
	reader:= bufio.NewReader(os.Stdin)
	x,_:= reader.ReadString('\n')
	y,_:=reader.ReadString('\n')
	sign,_:=reader.ReadString('\n')
	calculate(strings.TrimSpace(x),strings.TrimSpace(y),strings.TrimSpace(sign))
}

func calculate(x,y,sign string){
	x1:=convertValueIntoFloat(x)
	y1:=convertValueIntoFloat(y)
	var result float64
	switch sign{
	case "+":
		result=add(x1,y1)
	case "-":
		result=sub(x1,y1)
	case "*":
		result=mul(x1,y1)
	case "/":
		result=div(x1,y1)
	default:
		fmt.Println("No output")
	}
	fmt.Printf("%.0f%s%.0f = %.0f\n",x1,sign,y1,result)
}

func add(value1, value2 float64)float64{
	return value1+value2
}

func sub(value1, value2 float64)float64{
	return value1-value2
}

func mul(value1, value2 float64)float64{
	return value1*value2
}

func div(value1, value2 float64)float64{
	if value2==0{
		fmt.Println("Zero division error")
		return 0
	}else{
	return value1/value2
	}
}

func convertValueIntoFloat(value string) float64{
	floatValue, err:= strconv.ParseFloat(strings.TrimSpace(value),64)
	checkError(err)
	return floatValue
}

func checkError(err error){
	if err!=nil{
		panic(err)
	}
}

