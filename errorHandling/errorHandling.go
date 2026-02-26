package main

import (
	"errors"
	"fmt"
	"math"
	// "path/filepath"
)

func calculateArea(radious float64) (float64, error) {
	if radious < 0 {
		return 0, errors.New("Radious cannot be less than 0")
	}

	return math.Pi * float64(radious*radious), nil

}

func main() {
	// files, err := filepath.Glob("[")
	// if err != nil {
	// 	if errors.Is(err, filepath.ErrBadPattern) {
	// 		fmt.Println("Bad pattern error:", err)
	// 		return
	// 	}
	// 	fmt.Println("Generic error:", err)
	// 	return
	// }
	// fmt.Println("matched files", files)

	area, err := calculateArea(-1)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(area)

}
