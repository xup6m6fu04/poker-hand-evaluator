package main

import (
	"fmt"
	"rainbow/q1"
	"rainbow/q2"
	"rainbow/q3"
	"rainbow/q4"
	"rainbow/q5"
	"rainbow/q6"
	"rainbow/q7"
)

func main() {
	// Question 1
	fmt.Println("=========== Question 1 Start ===========")
	v1 := q1.Run()
	fmt.Println(v1)
	fmt.Println("=========== Question 1 End =============")

	// Question 2
	fmt.Println("=========== Question 2 Start ===========")
	v2 := q2.Run([]int{3, 4, 7, 7, 22, 9, 1})
	fmt.Println(v2)
	fmt.Println("=========== Question 2 End =============")

	// Question 3
	fmt.Println("=========== Question 3 Start ===========")
	v3 := q3.Run([]int{0x40a, 0x10b, 0x20c, 0x40d, 0x40e})
	fmt.Println(v3)
	fmt.Println("=========== Question 3 End =============")

	// Question 4
	fmt.Println("=========== Question 4 Start ===========")
	v4 := q4.Run([]int{0x308, 0x302, 0x209, 0x104, 0x406})
	fmt.Println(v4)
	fmt.Println("=========== Question 4 End =============")

	// Question 5
	fmt.Println("=========== Question 5 Start ===========")
	v5 := q5.Run([]int{0x10c, 0x30c}, []int{0x20c, 0x40b, 0x40c, 0x40d, 0x40e})
	fmt.Println(v5)
	fmt.Println("=========== Question 5 End =============")

	// Question 6
	fmt.Println("=========== Question 6 Start ===========")
	v6 := q6.Run([]int{0x406, 0x407}, []int{0x40a, 0x40b, 0x40c, 0x409, 0x408})
	fmt.Println(v6)
	fmt.Println("=========== Question 6 End =============")

	// Question 7
	fmt.Println("=========== Question 7 Start ===========")
	v7 := q7.Run([]int{0x50f, 0x202}, []int{0x105, 0x309, 0x40a, 0x40b, 0x40c})
	fmt.Println(v7)
	fmt.Println("=========== Question 7 End =============")
}
