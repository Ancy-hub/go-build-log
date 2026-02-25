package main

import (
	"fmt"
	"runtime"
)

/*given a large slice, split  into 4 parts, use go routines, sum each part concurrently, combine results. Hint: go routines and channels

*/

 func main(){
	var input=[]int{1,2,3,4,5}
	fmt.Print(concurrentSum(input))
 }

func sequentialSum(slice []int) int {
    sum := 0
    for _, v := range slice {
        sum += v
    }
    return sum
}

func concurrentSum(input []int)int{
	if len(input)<1000{
		return sequentialSum(input)
	}
	numWorkers:=runtime.NumCPU()
	ch:=make(chan int,numWorkers)
	partSize:=len(input)/numWorkers
	for i:=0;i<numWorkers;i++{
		start:=i*partSize
		end:=start+partSize
		if i== (numWorkers-1) {end=len(input)}

		sum:=0
		go func(part []int){
			for _,v:=range part{
				sum+=v
			}
		}(input[start:end])
		ch<-sum
	}

	total:=0
	for i:=0;i<numWorkers;i++{
		total+=<-ch
	}
	return total

}