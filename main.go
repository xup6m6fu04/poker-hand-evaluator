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
	v1 := q1.Run([]int{0x40a, 0x10b, 0x20c, 0x40d, 0x40e})
	fmt.Println(v1)
	fmt.Println("=========== Question 1 End =============")

	// Question 2
	fmt.Println("=========== Question 2 Start ===========")
	v2 := q2.Run([]int{0x308, 0x302, 0x209, 0x104, 0x406})
	fmt.Println(v2)
	fmt.Println("=========== Question 2 End =============")

	// Question 3
	fmt.Println("=========== Question 3 Start ===========")
	v3 := q3.Run([]int{0x10c, 0x30c}, []int{0x20c, 0x40b, 0x40c, 0x40d, 0x40e})
	fmt.Println(v3)
	fmt.Println("=========== Question 3 End =============")

	// Question 4
	fmt.Println("=========== Question 4 Start ===========")
	v4 := q4.Run([]int{0x406, 0x407}, []int{0x40a, 0x40b, 0x40c, 0x409, 0x408})
	fmt.Println(v4)
	fmt.Println("=========== Question 4 End =============")

	// Question 5
	fmt.Println("=========== Question 5 Start ===========")
	v5 := q5.Run([]int{0x50f, 0x202}, []int{0x105, 0x309, 0x40a, 0x40b, 0x40c})
	fmt.Println(v5)
	fmt.Println("=========== Question 5 End =============")
}
