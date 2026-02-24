// maps  := A map is a inbuilt data type in Go which is used to store key-value pairs

// make(map[type of key]type of value)

package main

import "fmt"

func main33() {
	currencyCode := make(map[string]string)

	// currencyCode1 := map[string]string {
	// 	"USD": "US Dollar",
	// 	"GBP": "Pound Sterling",
	// 	"EUR": "Euro",
	// }
	// //second way of creating map during declaration

	// fmt.Println(currencyCode)  //map[]

	currencyCode["USD"] = "US Dollar"
	currencyCode["GBP"] = "Pound Sterling"
	currencyCode["EUR"] = "Euro"
	currencyCode["INR"] = "Indian Rupee"

	// fmt.Println(currencyCode)
	// fmt.Println(currencyCode["INR"])

	// The zero value of a map is nil. If you try to add elements to a nil map, a run-time panic will occur. Hence the map has to be initialized before adding elements.

	var currencyCode2 map[string]string
	// currencyCode2["USD"] = "Dollar" //panic: assignment to entry in nil map  ***
	fmt.Println(currencyCode2)

	student := map[string]string{
		"name":   "Go",
		"age":    "2007",
		"syntax": "complex than python",
	}

	// fmt.Print(student ,"\n")
	// fmt.Println(student)
	// value, ok := map[key]

	// keyname :="name"

	// if keyname, ok := student[keyname]; ok {
	// 	fmt.Println( " Key Found ==>  ", keyname)
	// 	// return keyname
	// }

	for key, val := range student {

		fmt.Println(key, " : ", val)
	}

	// delete(student ,"name")  ***
	// fmt.Println(student)

	// Similar to slices, maps are reference types. When a map is assigned to a new variable, they both point to the same underlying data structure. Hence changes made in one will reflect in the other.

	// structs  := It is modified data type in which data can be grouped.

	// type currency struct {
	//     name string
	//     symbol  string
	// }

	// curUSA := currency{
	//     name :"Dollar",
	//     symbol :"$",
	// }
	// curInd := currency{
	// 	name:   "Rupess",
	// 	symbol: "@",
	// }

	// cocode := map [string] currency{
	//     "India" :  curInd,
	//     "USA" : curUSA ,
	// }

	// fmt.Println(cocode)

	// for key , value := range cocode{
	//     fmt.Printf("Contry name is %s and currency name is  %s and currency symbol is  %s \n" ,
	//     key ,value.name , value.symbol)
	// }

	type continents struct {
		name            string
		countriesCount  int
		peoplesColor    string
		rankofContinent int
	}

	asia := continents{
		name:            "Asia",
		countriesCount:  10,
		peoplesColor:    "brown",
		rankofContinent: 1,
	}
	america := continents{
		name:            "America",
		countriesCount:  54,
		peoplesColor:    "white",
		rankofContinent: 7,
	}
	europe := continents{
		name:            "Europe",
		countriesCount:  76,
		peoplesColor:    "white",
		rankofContinent: 4,
	}
	australia := continents{
		name:            "australia",
		countriesCount:  10,
		peoplesColor:    "white",
		rankofContinent: 5,
	}

	contRank := map[string]continents{
		"aus":      australia,
		"asia":     asia,
		"europe":   europe,
		"americas": america,
	}

	for key, val := range contRank {
		fmt.Printf("\n %s is %d th biggest country and has  %d countries and the peoples color is %s and short name i have used is %s", val.name, val.rankofContinent, val.countriesCount, val.peoplesColor, key)
	}

	// Maps can’t be compared using the == operator. The == can be only used to check if a map is nil.

}
