package main

import (
	"bufio"
	"calculatorpackage/calc"
	"fmt"
	"os"
	"strings"
)

/*
step: 1  Project Structure Setup
calculatorpackage/
├── main.go
└── calc/
    └── operations.go
step 2: create custom package (package name should match directory name. for eg: here in operations.go, I did package calc, because its inside calc directory)
step 3: Initialize go module. (Veryyyy imp):
cd calculatorpackage
go mod init calculatorpackage (always initialize inside project root folder)
This tells Go: "When you see calculatorpackage/calc in imports, look in the calc/ directory of this project."
Without it, Go has no way to map "calculatorpackage/calc" to your local calc/ folder. else it will search in go's offical packages
step 4: import path in main file: calculatorpackage/calc = module name + package directory

*/

func main() {
	fmt.Println("Calculator")
	reader := bufio.NewReader(os.Stdin)
	x, _ := reader.ReadString('\n')
	y, _ := reader.ReadString('\n')
	sign, _ := reader.ReadString('\n')
	calc.Calculate(strings.TrimSpace(x), strings.TrimSpace(y), strings.TrimSpace(sign))
}
