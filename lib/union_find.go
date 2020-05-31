package main

import "fmt"

func union(comps, size []int, i , j int) {
    ri := findRoot(comps, i)
    rj := findRoot(comps, j)
    if ri != rj {
        if size[ri] > size[rj] {
            comps[rj] = ri
            size[ri] += size[rj]
        } else {
            comps[ri] = rj
            size[rj] += size[ri]
        }
    }
}

func findRoot(comps []int, i int) int {
    if comps[i] == i {
        return i
    }
    res := findRoot(comps, comps[i])
    comps[i] = res
    return res
}

func main() {
	N := 10
	components := make([]int, N)
	size := make([]int, N)
	i := 0
	for ; i < n+1; i++ {
		components[i] = i
		size[i] = 1
	}
	union(components, size, 0, 2)
	union(components, size, 1, 4)
	union(components, size, 5, 4)
	union(components, size, 2, 5)
	union(components, size, 0, 2)
	fmt.Println(findRoot(components,5))
}
