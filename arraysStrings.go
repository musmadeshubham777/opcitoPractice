package main

import "fmt"

// pass by value not pass by reference

//

func changelocal(a [5]int) [5]int {
	fmt.Println("inside function")
	a[0] = 555 //this will not change original a means pass by value
	var b [5]int = a
	return b
}

func main22() {

	// a := [...] int {1,2,3,4,5}
	// fmt.Println("before calling function",a)
	// var result[5]int = changelocal(a)
	// fmt.Println("after calling function" , result)

	// fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	// fruitslice := fruitarray[1:3]
	// fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice))

	// //length of fruitslice is 2 and capacity is 6

	// countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	// neededCountries := countries[:len(countries)-2]
	// countriesCpy := make([]string, len(neededCountries))
	// copy(countriesCpy, neededCountries)
	// fmt.Println()
	// fmt.Println("Required Countries" , countriesCpy)

	// find(84 ,45,45,43,45,84)
	// change("helloo" ,"world")

	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}

func find(num int, nums ...int) {
	found := false

	for i, v := range nums {
		fmt.Println(i, v)
		if v == num {
			found = true
			println("Number Found ", v)
		}
	}

	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")

}

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)

}
