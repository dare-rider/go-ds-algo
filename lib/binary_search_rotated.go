package main

import "fmt"

func rotatedBinSearch(arr []int, ele int, l int, r int) int {
		if l > r {
			return -1
		}
		mid := (l + r)/2
		if arr[mid] == ele {
			return mid
		}
		if arr[mid] >= arr[l] {
			if ele <= arr[mid] && ele >= arr[l]{
				return rotatedBinSearch(arr, ele, l, mid -1)
			}
			return rotatedBinSearch(arr, ele, mid + 1, r)
		} else {
			if ele >= arr[mid] && ele <= arr[r]{
				return rotatedBinSearch(arr, ele, mid + 1, r)
			}
			return rotatedBinSearch(arr, ele, l, mid -1)
		}
}


func main(){
	arr := []int{2,3,4,5,1}
	fmt.Println(rotatedBinSearch(arr, 1, 0 , len(arr) - 1))
}