package main

import "fmt"

// "fmt"
// "learnpackage/simpleinterest" //importing custom package
// "log"

var p, r, t = 5000.0, 10.0, 1.0

/*
* init function to check if p, r and t are greater than zero
 */

// func init() {
// fmt.Println("Main package initialized")
// if p < 0 {
// 	log.Fatal("Principal is less than zero")
// }
// if r < 0 {
// 	log.Fatal("Rate of interest is less than zero")
// }
// if t < 0 {
// 	log.Fatal("Duration is less than zero")
// }else{
// 	log.Fatal("Duration is zero")
// }
// }

func main12() {
	// fmt.Println("Simple interest calculation")
	// // si := simpleinterest.Calculate(p, r, t)
	// fmt.Println("Simple interest is", 5000)

	// for i := 0; i < 5; i++ {
	// 	for j :=0 ; j<=i ; j++{
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// switch case

	// var key int =10 ;
	// switch key{
	// case 1:
	// 	fmt.Println("case 1 ")
	// case 2:
	// 	fmt.Println("case 2 ")
	// case 3:
	// 	fmt.Println("case 3 ")
	// case 10:
	// 	fmt.Println("case 10 ")
	// default :
	// 	fmt.Println("this is default case")
	// }

	//case cannot be duplicated

	// hour := 15
	// Using switch to determine the work shift
	// switch {
	// case hour >= 6 && hour < 12:
	// }

	// arrays

	// var a[5] int   //normal array declaration

	// a := [5]int {1,2,3,4,5} //short hand operator to declare array.==> [1,2,3,4,5]
	// b := [5]int {1,2} //short hand operator to declare array.==> [1,2,0,0,0]
	// fmt.Println(a)
	// fmt.Println(b)

	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)

}
