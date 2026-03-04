package main

//create slice of student structs, convert to json, save to a file, read back from file, print data. hint: encoding/json
import (
	"encoding/json"
	"fmt"
	"os"
)

/*
In Go, you use the encoding/json package to convert data to JSON. Here are the key functions:
Convert to JSON:
json.Marshal(v interface{}) ([]byte, error) - converts Go data to JSON bytes
json.MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) - converts with formatting
Convert from JSON:
json.Unmarshal(data []byte, v interface{}) error - converts JSON bytes to Go data
In Go, you write to files using the os.WriteFile("filename.txt", data, 0644). The 0644 is the file permission (readable by all, writable by owner).
If file already exists then just file.Write(data)
*/
func main() {
	students := []student{
		{Name: "Ancy", Age: 23},
		{Name: "Gouri", Age: 23},
	}

	data, _ := json.Marshal(students)
	os.WriteFile("studentdetails.json", data, 0644)

	fileData, _ := os.ReadFile("studentdetails.json")
	var s []student
	json.Unmarshal(fileData, &s)
	fmt.Print(s)
}

type student struct {
	Name string
	Age  int
}
