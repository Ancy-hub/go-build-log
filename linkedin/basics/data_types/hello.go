package main
import(
	"fmt"
)

func main(){

	str1:= "Hi"

	fmt.Println(str1)

	fmt.Println("Hello world")

	stringLenght, err:= fmt.Println("Hi world")
	if(err==nil){
		fmt.Println("String lenght:", stringLenght)
	}

	fmt.Printf("Value of number: %v\n", 100)
	//v is value, T is type.
	fmt.Printf("Datat type: %T\n",100)
}