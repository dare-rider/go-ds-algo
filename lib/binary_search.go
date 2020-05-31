package main

import "fmt"

func binSearch(arr []int, ele int) int {
	l := 0
	r := len(arr) - 1
	for {
		if l > r {
			break
		}
		mid := (l + r)/2
		if arr[mid] == ele {
			return mid
		} else if arr[mid] < ele {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}


func main(){
	arr := []int{1,2,3,4,5}
	fmt.Println(binSearch(arr, 4))
}
