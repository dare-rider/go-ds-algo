package main

//There're 3 kinds of the bus ticket. 1: ticket 1 cost 2 and can be used for a day.
// 2: ticket 2 cost 7 and can be used for a consecutive 7 days. 3: ticket 3 cost 25 can be used for a month.
// Assume month here means 30 consecutive days.
//Now there's an array filled with elements.
// Each element value is a date for a person to travel.
// This array has already been sorted in ascending order, like {1,3,4,5,11,12,23,24,30}.
// The final day is 30th and the first day is 1st.

import "fmt"

func main() {
	arr := []int{1,3,4,5,11,12}
	dp := make([]int, len(arr))
	cst := costCal(arr, 0, 0, -1, dp)
	if cst > 25{
		fmt.Println(25)
	} else {
		fmt.Println(cst)
	}
	fmt.Println(dp)
}

func costCal(arr []int, idx int, cost int, sStart int, dp []int) int{
	if idx == len(arr) {
		return cost
	}
	var c1 int
	var c2 int
	if idx == 0 || sStart == -1 {
		c1 = costCal(arr, idx + 1, cost + 2, -1, dp)
		c2 = costCal(arr, idx + 1, cost + 7, arr[idx], dp)
	} else {
		c1 = costCal(arr, idx + 1, cost + 2, -1, dp)
		if (arr[idx] - sStart < 7) {
			c2 = costCal(arr, idx + 1, cost, sStart, dp)
		} else {
			c2 = costCal(arr, idx + 1, cost + 7, arr[idx], dp)
		}
	}
	var mincost int
	if c1 < c2 {
		mincost =  c1
	} else {
		mincost =  c2
	}
	return mincost
}

