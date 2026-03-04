package main
//give a slice, return a new slice without duplicates. maps,slice building

import(
	"fmt"
	"sort"
)

func main(){
	var given=[...]int{1,2,3,4,1,2,3,5,6,7}
	count:=make(map[int]int,len(given))
	for _,number:=range given{
		count[number]++
	}
	//fmt.Println(count)
	var output []int
	for k:=range count{
		output=append(output,k)
	}
	sort.Ints(output)

	fmt.Print(output)
}