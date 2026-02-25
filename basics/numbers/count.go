package main

import(
	"fmt"
)

func main(){
	var count int
	fmt.Scan(&count)
	numbers := make([]int, count)
	for i := range numbers {
		fmt.Scan(&numbers[i])
	}
	fmt.Println(numbers)
	var pos,neg,zero int
	for i:= range numbers{
		if numbers[i]>0{
			pos+=1
		}else if numbers[i]<0{
			neg+=1
		}else{
			zero+=1
		}
	}
	fmt.Printf("+ve digits: %v, -ve digits: %v, zeroes: %v",pos,neg,zero)
	
}