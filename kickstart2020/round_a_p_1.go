package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution(arr []int, size, money int) int{
	dpSum := make([]int, size)
	dpCnt := make([]int, size)
	maxCnt := 0
	for i := range arr {
		dpSum[i] = arr[i]
		dpCnt[i] = 1
		if arr[i] < money {
			maxCnt = 1
		}
	}
	for j := 1; j < len(arr); j++ {
		for i := 0; i < j; i++ {
			if dpSum[i] + arr[j] <= money && dpCnt[i] + 1 > dpCnt[j]{
				dpCnt[j] = dpCnt[i] + 1
				dpSum[j] = dpSum[i] + arr[j]
				if dpCnt[j] > maxCnt {
					maxCnt = dpCnt[j]
				}
			}
		}
	}
	return maxCnt
}

func main() {
	var tcCnt, size, money int
	reader := bufio.NewReader(os.Stdin)
	fmt.Scanf("%d", &tcCnt)
	for i := 1; i <= tcCnt; i++ {
		var inpArr []int
		line1, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		line1Inp := strings.Split(strings.TrimSuffix(line1, "\n"), " ")
		size, _ = strconv.Atoi(line1Inp[0])
		money, _ = strconv.Atoi(line1Inp[1])
		inpArrStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		inpArrStr = strings.TrimSuffix(inpArrStr, "\n")
		for _, s := range strings.Split(inpArrStr, " ") {
			v, _ := strconv.Atoi(s)
			inpArr = append(inpArr, v)
		}
		fmt.Printf("Case #%d: %d\n", i, solution(inpArr, size, money))
	}
	
}
