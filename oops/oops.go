package main

import (
	"fmt"
	// "go/types/"
)

//Go is not a pure object oriented programming language.

//composition means embeding one struct into other

type author struct {
	firstName string
	lastName  string
	bio       string
}

type blogPost struct {
	title   string
	content string
	author
}

func (w blogPost) display() {
	fmt.Println(w.author)
	fmt.Println(w.content)
}

type website struct {
	blogPosts []blogPost
}

type vlogger struct {
	blogposts []blogPost
}

func (b website) describe() {

	fmt.Println("Printing Blogs")

	for _, v := range b.blogPosts {
		v.display()
		fmt.Println()
	}

}

func mainoops() {

	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	blogPost1 := blogPost{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	blogPost2 := blogPost{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	blogPost3 := blogPost{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}

	w := website{
		blogPosts: []blogPost{blogPost1, blogPost2, blogPost3},
	}
	w.describe()

}

// type Vehicle struct{
// 	Type string
// 	wheels int
// }

// type Car struct{
// 	Vehicle
// 	Name string
// 	Fueltype string
// }

// func main() {
// 	c := Car{}
// 	c.Type="SedanCar"
// 	c.wheels=4
// 	c.Name="InC"
// 	c.Fueltype ="Disel"
// 	fmt.Println(c)
// }
