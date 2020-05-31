package main

import "fmt"

func countRemoval(str string, l int, r int, dp [][]int) int{
	if l > r || l == r {
		return 0
	}
	if r == l + 1 {
		if str[l] == str[r] {
			return 0
		} else {
			return 1
		}
	}
	if dp[l][r] != 0 {
		return dp[l][r]
	}
	var cnt int
	if str[l]  == str[r] {
		cnt = countRemoval(str, l + 1, r - 1, dp)
	} else {
		cnt = 1 + countRemoval(str, l + 1, r, dp)
		cnt2 := 1 + countRemoval(str, l, r - 1, dp)
		cnt3 := 2 + countRemoval(str, l + 1, r - 1, dp)
		if cnt2 < cnt {
			cnt = cnt2
		}
		if cnt3 < cnt {
			cnt = cnt3
		}
	}
	dp[l][r] = cnt
	return cnt
}

func main() {
	str := "aaabbbbaba"
	dp := make([][]int, len(str))
	for i, _ := range dp {
		dp[i] = make([]int, len(str))
	}
	fmt.Println(countRemoval(str, 0, len(str) - 1, dp))
	fmt.Println(dp)
}
