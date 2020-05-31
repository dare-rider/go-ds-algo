package main

import "fmt"

// tries to find the length og such prefix which is also a suffix
// eg: ACAC has such prefix `AC`
// ACACA has such prefix `ACA`
// ACAGT has such prefix ``

//complexities:
//	m = len(pat)
//	n = len(str)
//	space = O(m)
//	time = O(m + n)

func prefixArray(in []int, pat string){
	l := 0
	i := 1
	in[l] = 0
	for i < len(pat) {
		if pat[l] == pat[i] {
			in[i] = l + 1
			i += 1
			l += 1
		} else {
			if l == 0 {
				in[i] = 0
				i += 1
			} else {
				l = in[l-1]
			}
		}
	}
}

func kmpSearch(prefix []int, str string, pat string) bool {
	i := 0
	j := 0
	var ind []int
	for i <= len(str) - len(pat) {
		for i + j <= len(str) && j < len(pat) && str[i + j] == pat[j] {
			j += 1
		}
		if j == len(pat) {
			ind = append(ind, i)
			i = i + len(pat)
			j = 0
		} else {
			if j > 0 {
				// ignoring the prefix which is already matched
				// don't bother on i, used to calculate the starting index of pattern
				i = i + j - prefix[j - 1]
				j = prefix[j - 1]
			} else {
				i += 1
			}
		}
	}
	fmt.Println(ind)
	return len(ind) > 0
}

func main() {
	//pat := "ACACAGTACC"
	//str := "DBSACACACACAGTACCIIO"
	pat := "BARBER"
	str := "THIS IS MEBAER BARBER"
	in := make([]int, len(pat))
	prefixArray(in, pat)
	fmt.Println(kmpSearch(in, str, pat))
}
