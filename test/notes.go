package main

import (
	// "fmt"
	// "math"
)

// three ways to run go file
//go run heelo.go
//go install
//go build
// Variable is the name given to a memory location to store a value of a specific type



func main1(){
	// var age int16 =25 ;//if i dint give default value it wil be 0
	// var price ,quantity int = 100 ,2;

	// fmt.Println("age is" , age)
	// age= 21
	// fmt.Println("age is" , age)
	// age=22
	// fmt.Println("age is" , age)
	// age=23-age
	// fmt.Println("age is" , age)
	// fmt.Println("final amount is" , price * quantity)


	// var (
	// 	name= "navib"
	// 	age1=38
	// 	height int = 174
	// )

	// fmt.Println("name is" , name  , "age is" , age1  ," and height is" ,height)

//short hand operator  :=

// var aa , bb int  = 20 , 30 ;

// aa ,cc  :=20 ,40 ;  // if i make bb instead of cc it will give error  //no new variables on left side of :=

// fmt.Println(aa, bb , cc )



//type casting

// var ab =10.00 ;
// var ac = 20

// var ar= ab + ac; //invalid operation: ab + ac (mismatched types int and string)
// var ar= float64(ab + float64(ac));
// fmt.Println(ar)



//constants

//you cannot assign value to the constants again if the variable  is already assigned with the value
// const zz= 25;
// zz=10; //this is not allowed to assign values to the constant again

// var a= math.Sqrt(4) //this is allowed with  variable  but not with the constants
// const b = math.Sqrt(4) //not allowed because value of constant should be known at compile time
// fmt.Println(b)

// var defaultName = "Sam" //allowed
// type myString string
// var customName myString = "Sam" //allowed
// customName = defaultName        //not allowed


// var defaultName = "Sam" ; //allowed
// type myString string ; 
// var customName myString = "Sam" ;    //allowed
// customName = myString(defaultName) ;


}

