//string
// Strings in Go are immutable because they are implemented as read-only byte slices. To modify them, we must convert them into a mutable slice like []byte or []rune.

package main

import "fmt"

func main44() {

	newString := "Hello Go Lang"
	count := len(newString)

	for i := 0; i < count; i++ {
		// fmt.Printf( "%c", newString[i] )  // 48 65 6c 6c 6f 20 47 6f 20 4c 61 6e 67
		// A string is a slice of bytes, it’s possible to access each byte of a string.
		// %c is used to print each character of string
		// %x is used to print byte code of each character of string
	}

	//reverse a string

	strngaaaa := "helooaaa"
	runes := []rune(strngaaaa)
	fmt.Print(runes)
	ln := len(strngaaaa) - 1

	for i, j := 0, ln; i < j; i, j = i+1, j-1 {
		// runes[i] , runes[j] = runes[j] ,runes[i]
	}

	// fmt.Println(strngaaaa)

	//rune  :::
	//An alias for int32

	//strings can be compared using ==
	// can be concatenated using +
	// result := fmt.Sprintf("%s %s", string1, string2)
	//any valid unicode character within single quote is a rune

}
