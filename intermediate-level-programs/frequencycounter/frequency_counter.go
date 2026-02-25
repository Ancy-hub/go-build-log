package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main(){
	reader:=bufio.NewReader(os.Stdin)
 	sentance,_:=reader.ReadString('\n')
	words:=strings.Fields(sentance)
	count:=make(map[string]int,0)
	
	for _, word := range words {
		count[word]++
	}
	
	keys:=make([]string,len(count))
	i:=0
	for k:= range count{
		keys[i]=k
		i++
	}

	sort.Strings(keys)
	for i:= range keys{
		fmt.Printf("%s: %d\n",keys[i],count[keys[i]])
	}

}