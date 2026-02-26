//pointers go

package main

import "fmt"



func mainpointer() {
	// fmt.Println("he/llo")

	// a := 255
	// b := &a //A pointer is a variable that stores the memory address of another variable. Here a is pointer of a
	// c:= b

	// fmt.Println("Value of a",a)
	// fmt.Println("Address of a",b)
	// fmt.Println("Value of   c",c)

	//Declaring Pointers

	// *T is the type of the pointer variable which points to a value of type T.
	// d := 255
	// var e *int = &d
	// fmt.Printf("Type of e is %T\n", e)
	// fmt.Println("address of d is", e)

	// The & operator is used to get the address of a variable

	//Zero Value Pointer

	// a:= 25
	// var b * int

	// if b == nil {
	//     fmt.Println("b is", b)
	// 	b = &a
	// 	fmt.Println("b after initialization is", b)
	// }

	//Creating pointers using new function

	// size := new(int)
	// fmt.Printf("The size of size %d and address is %v and the type of size is %T" ,*size ,size ,size )

	// a := 25
	// b := &a
	// *b++  //this will increment value of a  because b  points to the address of the a
	// fmt.Printf("%T", a)
	// fmt.Println()

	// fmt.Printf("%T" , b)

	// fmt.Print(a)

	// fmt.Println(a)  //this also give value of a
	// fmt.Println(*b)  //this will give value of a

    //passing pointer to the function 

    // aa :=58
    // fmt.Println("before calling value" , aa)

    // zval := &aa  //instead of passing value passed pointer to the function.
    // changes(zval)
    // fmt.Println("value of a after function call is", aa)

    //    r := return_Pointer()
    //    fmt.Println(r)

    as:= [3]int{100,200,300}
    fmt.Println(as)

    result :=modArrar(as)
    fmt.Println(result)

    // Go does not support pointer arithmetic which is present in other languages like C and C++.

}



//returning pointer from function

func return_Pointer() *int {
    a := 24
    return &a
}

func modArrar (lss [3] int ) [3] int {
    lss[0] =90
    return lss
}


func changes(val *int) {
	*val = 55
}





