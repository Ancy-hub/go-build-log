package main

//add emp, search emp by id, find highest and lowest sal, find avg salary
import "fmt"

func main() {
	var count int
	fmt.Print("No of emps: ")
	fmt.Scanln(&count)
	org := make([]Emp, count)
	for i := range count {
		fmt.Print("Enter emp id: ")
		fmt.Scanln(&org[i].Id)
		fmt.Print("Enter emp name: ")
		fmt.Scanln(&org[i].Name)
		fmt.Print("Enter emp salary: ")
		fmt.Scanln(&org[i].Salary)
	}

	var findID int
	fmt.Print("Enter emp id to find: ")
	fmt.Scanln(&findID)

	high, low, total := org[0].Salary, org[0].Salary, 0.0
	found := false
	for _, person := range org {
		if person.Id == findID {
			fmt.Println(person)
			found = true
		}
		if person.Salary > high {
			high = person.Salary
		}
		if person.Salary < low {
			low = person.Salary
		}
		total += person.Salary
	}

	if !found {
		fmt.Println("Employee not found")
	}

	avg := total / float64(count)
	fmt.Printf("Highest: %.1f, Lowest: %.1f, Avg: %.1f\n", high, low, avg)
}

type Emp struct {
	Id     int
	Name   string
	Salary float64
}

/*Actually, you're right that the else if approach works correctly for min/max finding! The separate if statements are more about code clarity and independence rather than fixing a bug.

The main benefit is logical independence - each condition represents a separate concern (finding max vs finding min), so they should be evaluated independently rather than being mutually exclusive.

Both approaches work correctly, but separate if statements are cleaner and more maintainable.*/
