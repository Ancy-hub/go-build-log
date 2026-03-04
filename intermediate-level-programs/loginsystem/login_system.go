package main

import "fmt"

func main() {
	login := map[string]string{
		"Ancy":  "Ancy@2002",
		"Gouri": "Gouri@1234",
	}

	var username, password string
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	if login[username] == password && username != "" {
		fmt.Println("Login successful")
	} else {
		fmt.Println("Invalid credentials")
	}
}

/*
func main(){
	login:=make(map[string]string)
	login["Ancy"]="Ancy@2002"
	login["Gouri"]="Gouri@1234"
	
	var username, password string
	fmt.Println("Enter you username: ")
	fmt.Scanln(&username)
	fmt.Println("Enter you password: ")
	fmt.Scanln(&password)

	if login[username]==password{
		fmt.Println("Login successfull")
	}else{
		fmt.Println("Invalid credentials")
	}

}*/