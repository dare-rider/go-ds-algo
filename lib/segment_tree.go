package main

import "fmt"

func prepareSegmentTree(input []int, sTree []int, lo, hi, pos int) {
	if lo == hi {
		sTree[pos] = input[lo]
		return
	}
	mid := (lo + hi)/2
	prepareSegmentTree(input, sTree, lo, mid, 2*pos + 1)
	prepareSegmentTree(input, sTree, mid + 1, hi, 2*pos + 2)
	min := sTree[2*pos + 1]
	if sTree[2*pos + 2] < min {
		min = sTree[2*pos + 2]
	}
	sTree[pos] = min
}

// full overlap return data on pos
// no overlap return max
// partial overlap - min (left child, right child)
func returnMinForRange(sTree []int, qLow, qHigh, lo, hi, pos int) int{
	if lo >= qLow && hi <= qHigh {
		return sTree[pos]
	}
	if qLow > hi || qHigh < lo {
		return 1 << 31
	}
	mid := (lo + hi)/2
	l := returnMinForRange(sTree, qLow, qHigh, lo, mid, 2*pos + 1)
	r := returnMinForRange(sTree, qLow, qHigh, mid + 1, hi, 2*pos + 2)
	min := l
	if r < min{
		min = r
	}
	return min
}

func main() {
	var input = []int{-1,3,0,2,5}
	sTreeSize := 1
	for sTreeSize < len(input) {
		sTreeSize = (2 * sTreeSize)
	}
	sTreeSize = 2 * sTreeSize - 1
	sTree := make([]int, sTreeSize)
	prepareSegmentTree(input, sTree, 0, len(input) - 1, 0)
	res := returnMinForRange(sTree, 3, 4, 0, len(input) - 1, 0)
	fmt.Println(res)
}
