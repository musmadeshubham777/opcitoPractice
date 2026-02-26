//methods -- A method is just a function with a special receiver type between the func keyword and the method name. The receiver can either be a struct type or non-struct type

// Generally, pointer receivers can be used when changes made to the receiver inside the method should be visible to the caller.

// Pointers receivers can also be used in places where it’s expensive to copy a data structure.
//  Consider a struct that has many fields. Using this struct as a value receiver in a method will need the entire struct to be copied which will be expensive. In this case, if a pointer receiver is used, the struct will not be copied and only a pointer to it will be used in the method.

// func (a int) add(b int) { //compilation error cannot define new methods on non-local type int
// }

// func main() {

// }  This is not allowed because defination of int and defination of add not in same package
// To make this working we need to make alias like
// type myint int now this will work

// type myint int
// func(i myint) test(y myint){
// }

package main

import "fmt"

// func (t Type) methodName(parameter list) {
// }

type employees struct {
	id    int
	fname string
	lname string
}

type car struct {
	id      int
	carType string
}

type rectangle struct {
	length int
	width  int
}

//this is value reciver
// func (emp employees) test() {
// fmt.Println(emp.id)
// fmt.Println(emp.fname)
// fmt.Println(emp.lname)

// }

//now lets have one example of pointer reciever.
// func (e *employees)displayEmployee() {

// fmt.Println(e.lname)

// }

// func displayE(e employees) {

// fmt.Println(e.fname, e.lname, e.id)
// }

// func displayCarUsingArgument(c car){
//     fmt.Println("displayCarUsingArgument" ,c.id , c.carType)
// }

// func (c car)displayCarUsingS(d car){
//     fmt.Println("displayCarUsingS",c.id , c.carType)
//     fmt.Println("displayCarUsingS",d.id , d.carType)
// }

// func (c *car)displayCarUsingPointer(){
//     fmt.Println("displayCarUsingPointer",c.id , c.carType)

// }

func perimeter(r *rectangle) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *rectangle) perimeter() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

func mainMethods() {
	// e1 := employees{
	// 	id:    14,
	// 	fname: "Shubham",
	// 	lname: "Musmade",
	// }

	// e1.test() // this way i can call function
	// displayE(e1)

	// adressofE1 := &e1
	// adressofE1.displayEmployee()

	// c1 :=car{ id:1 , carType : "sedan" }
	// displayCarUsingArgument(c1)
	// c1.displayCarUsingS(c1)
	// c1.displayCarUsingPointer()

	// re1 :=rectangle{
	//     length: 10,
	//     width: 20,
	// }

	r := rectangle{
		length: 10,
		width:  5,
	}
	p := &r //pointer to r
	perimeter(p)
	p.perimeter()

	/* cannot use r (type rectangle) as type *rectangle in argument to perimeter*/

	// perimeter(r)

	r.perimeter() //calling pointer receiver with a value

	// a :=10
	// result := a.add(20)

	var a myint = 10
	result := a.add(20) // 20 automatically converted to myint
	fmt.Println(result)

}

type myint int

func (a myint) add(b myint) myint {
	return a + b
}
