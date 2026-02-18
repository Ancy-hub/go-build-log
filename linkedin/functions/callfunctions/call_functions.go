package main
//go is organized with packages and packages have functions.
//functions within same package can be named anything, initial character can be upper or lowercase. 
//if initial character is lowercase, then only available within the current package. if its uppercase, then its public to the application.
import "fmt"

func main(){
	fmt.Println("Functions in go")
	doSomething()
	fmt.Println(addValues(2,3,4))
}

func doSomething(){
	fmt.Println("Doing something")
}

func add(x, y, z int) int{
	v:=x+y+z
	return v
}

func addValues(values ...int) (int, int , float64) {
	sum:=0
	for _,v:= range values{
		sum+=v
	}

	count:=len(values)
	avg:=float64(sum)/float64(count)
	return sum,count,avg
}