package main

/*Go has set of tools to work easily with web service apis.
The http package lets you make request and send data to remote hosts, it also lets you create HTTP server applications that listen for and respond to requests.

*/

import (
	"fmt"
	"io"
	"net/http"
)

const url="http://services.explorecalifornia.org/json/tours.php"

func main(){
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
	fmt.Print(content)
	
}

func checkError(err error){
	if err!=nil{
		panic(err)
	}
}
