package main

import (
	"fmt"
	"strings"
	"sort"
	)

func main() {
	arr := []int{-2,-3,4,-1,-2,1,5,-3}
	l, r, sum := maxSum(arr)
	fmt.Println(l,r,sum)
	strings.NewReader("hello")
}

func maxSum(arr []int) (int, int, int){
	l := 0
	r := 0
	sum := 0
	sumTillNow := 0
	for i, ele := range arr {
		sumTillNow += ele
		if sumTillNow < 0 {
			sumTillNow = 0
			l = i + 1
		}
		if sumTillNow > sum {
			sum = sumTillNow
			r = i
		}
	}
	return l,r,sum
}