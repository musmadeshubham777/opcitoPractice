package main

import (
	"fmt"
	// "go/types"
)

func mainStruct() {

	// fmt.Println("Entered")
	structStd()

}

func structStd() {
	// main()
	fmt.Println("Entered structStd")

	type employee struct {
		name        string
		age         int
		department  string
		reportingTo string
	}

	emp1 := employee{"shubham", 28, "Core Engineering", "Manoj"}
	emp2 := employee{"vaibhav", 25, "Core Engineering", "Yograj"}
	emp3 := employee{"Shivraj", 25, "Core Engineering", "Yograj"}
	emp4 := employee{"Test", 32, "Core Engineering", "Test"}
	var emp5 employee //zero valur struct

	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp3)
	fmt.Println(emp4)
	fmt.Println(emp5)

	// It is also possible to specify values for some fields and ignore the rest. In this case, the ignored fields are assigned zero values.

	emp7 := &employee{
		name: "Test",
		age:  32}
	fmt.Println(*emp7)

}
