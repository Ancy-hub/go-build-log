package main

/*GO has a json package that helps you easily create and read JSON formatted text.
 */

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const url="http://services.explorecalifornia.org/json/tours.php"

func main(){
	content:=readHttpContent()
	//fmt.Print(content)
	tours:= toursFromJSON(content)
	for _, tour:= range tours{
		price,_:= strconv.ParseFloat(tour.Price,16)
		// $ sign is a literal string
		fmt.Printf("%v: ($%v)\n", tour.Name, price)
	}
}

func toursFromJSON(content string) []Tour{
	tours:=make([]Tour,0)
	decoder:= json.NewDecoder(strings.NewReader(content))
	//most imp thing abt this Token fn does is to guarentee that the delimiters,  the commas and braces are properly nested and matched
	_ ,err:= decoder.Token()
	checkError(err)

	var tour Tour
	for decoder.More(){
		err:=decoder.Decode(&tour)
		checkError(err)
		tours=append(tours,tour)
	}
	return tours

}

func readHttpContent() string{
		//while creating a object, we can pass a whole bunch of parameters, not passing any means accepting the default value.
		client:=http.Client{}
		req, err:= http.NewRequest("GET",url, nil)
		checkError(err)
		//we might get a access forbidden message. it means the this particular api needs something called an user agent
		req.Header.Set("User-Agent","")
	
		resp, err:= client.Do(req)
		checkError(err)
		defer resp.Body.Close()
	
		fmt.Printf("Response type: %T\n", resp)
		bytes, err:= io.ReadAll(resp.Body)
		checkError(err)
		content:= string(bytes)
		return content
}

func checkError(err error){
	if err!=nil{
		panic(err)
	}
}

type Tour struct{
	Name, Price string
}