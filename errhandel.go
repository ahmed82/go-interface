package main

import (
	"fmt"
	"math"
)

func errHandel() {
	x := -5.0
	i, err := sql(x)
	if err != nil {
		//fmt.Println("There is an error")
		fmt.Println(err)
	}

	fmt.Println(i)
}

func sql(x float64) (float64, error) {
	if x < 0 {
		//return 0, errors.New("X cannot be less than 0")
		return 0, fmt.Errorf("the value of x %f cannot be 0", x)
	}
	return math.Sqrt(x), nil

}
