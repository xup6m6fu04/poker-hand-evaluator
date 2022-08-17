package main

import (
	"fmt"
	"poker-hand-evaluator/q1"
	"poker-hand-evaluator/q2"
	"poker-hand-evaluator/q3"
	"poker-hand-evaluator/q4"
	"poker-hand-evaluator/q5"
)

func main() {
	// Question 1
	fmt.Println("=========== Question 1 Start ===========")
	v1 := q1.Run([]string{"2d", "3d", "4s", "5d", "6d"})
	fmt.Println(v1)
	fmt.Println("=========== Question 1 End =============")

	// Question 2
	fmt.Println("=========== Question 2 Start ===========")
	v2 := q2.Run([]string{"2d", "Js", "4d", "5d", "6d"})
	fmt.Println(v2)
	fmt.Println("=========== Question 2 End =============")

	// Question 3
	fmt.Println("=========== Question 3 Start ===========")
	v3 := q3.Run([]string{"2d", "3d"}, []string{"4d", "5d", "6d", "7d", "8d"})
	fmt.Println(v3)
	fmt.Println("=========== Question 3 End =============")

	// Question 4
	fmt.Println("=========== Question 4 Start ===========")
	v4 := q4.Run([]string{"Ts", "Js"}, []string{"Qs", "Ks", "As", "Td", "Kd"})
	fmt.Println(v4)
	fmt.Println("=========== Question 4 End =============")

	// Question 5
	fmt.Println("=========== Question 5 Start ===========")
	v5 := q5.Run([]string{"2s", "3c"}, []string{"BJ", "Td", "6c", "7d", "8d"})
	fmt.Println(v5)
	fmt.Println("=========== Question 5 End =============")
}
