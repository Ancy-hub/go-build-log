package main

import(
	"fmt"
	"sort"
)

//Map is unordered collection of key vaue pairs. its a hash table that lets you store collections of data and then find them in collection by their keys.
//Map's key can be of any type thats comparable for the purposes of sorting, commonly strings are used.


func main(){
	m:= make(map[string]int)//menas that keys are strings and associated values are integers. always initialise maps with make() fn.
	m["Key"]=42
	fmt.Println(m)
	
	colors:=make(map[string]string)
	colors["sky"]="Blue"
	colors["sea"]="Green"
	fmt.Println(colors)
	Blue:=colors["sky"]
	fmt.Println(Blue)

	// delete(colors,"sky")
	// fmt.Println(colors)

	for k,v:=range colors{
		fmt.Printf("%v: %v\n",k,v)
	}
	//to print a map in the alphabetical order, extract the keys from the map as a slice of strings.
	keys:=make([]string,len(colors))
	i:=0
	for k:=range colors{
		keys[i]=k
		i++
	}

	sort.Strings(keys)
	for i:=range keys{
		fmt.Println(colors[keys[i]])
	}

}