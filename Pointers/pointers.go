package main

import(
	"fmt"
)

func main(){
	age := 25
	agePointer := &age 
	fmt.Println("Age:", age)
	fmt.Println("Age Pointer:", agePointer)
	fmt.Println("Adult Years:", getAdultYears(agePointer))
}

func getAdultYears(age *int) int{
	return *age - 18
}