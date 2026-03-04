package main

import "fmt"

func main(){
	var count int
	fmt.Print("No. of students: ")
	fmt.Scanln(&count)
	students := make([]Student, count)
	for i := range count{
		fmt.Print("Enter student name: ")
		fmt.Scanln(&students[i].Name)
		fmt.Print("Enter maths mark: ")
		fmt.Scanln(&students[i].Math)
		fmt.Print("Enter science mark: ")
		fmt.Scanln(&students[i].Science)
		fmt.Print("Enter English mark: ")
		fmt.Scanln(&students[i].English)
	}

	for _, s := range students{
		total := s.Math + s.Science + s.English
		fmt.Printf("For %s: highest=%.1f, total=%.1f, average=%.1f\n", s.Name, max(s.Math, s.Science, s.English), total, total/3)
	}
}

type Student struct{
	Name string
	Math, Science, English float64
}

// func (s Student) average() float64{
// 	avg:=(s.English+s.Math+s.Science)/float64(3)
// 	return math.Round(avg)/100
// }
// func (s Student) total() float64{
// 	total:=s.English+s.Math+s.Science
// 	return total
// }
// func (s Student) highest()float64{
// 	highest:=max(s.English,s.Math,s.Science)
// 	return highest
// }

