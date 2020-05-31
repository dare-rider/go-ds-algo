package main

import "fmt"

//BM based on 2 subparts
//* bad character
//* good suffix (trivial to implement, ignoring) thus called boyer moore horspool algo
//- comparision start from right to left

// boyer moore horspool algo complexities (considering bad character only)
// Complexities, pattern length m, string length n
// Space O(m)
// Time Best O(m + n), worst O(mn)

func prepareMap(mp map[byte]int, pat string) {
	for i := range pat {
		// ignoring last char in pattern
		if i < len(pat) - 1 {
			mp[pat[i]] = len(pat) - i - 1
		}
	}
}

func bmSearch(mp map[byte]int, str string, pat string) bool {
	i := 0
	j := len(pat) - 1
	var ind []int
	for i + j < len(str) {
		for i + j < len(str) && j >= 0 && str[i + j] == pat[j] {
			j -= 1
		}
		if j < 0 {
			ind = append(ind, i)
			i = i + len(pat)
		} else {
			if _, ok := mp[str[i + j]]; !ok {
				i = i + j + 1
			} else {
				// last index - j
				lenSuccessfullyTraced := len(pat) - 1 - j
				i = i - lenSuccessfullyTraced + mp[str[i + j]]
			}
		}
		j = len(pat) - 1
	}
	fmt.Println(ind)
	return len(ind) > 0
}

func main() {
	pat := "BARBER"
	str := "THIS IS MEBAER BARBER"
	//pat := "ACACAGTACC"
	//str := "DBSACACACACAGTACCIIO"
	badCharShiftMap := make(map[byte]int)
	prepareMap(badCharShiftMap, pat)
	fmt.Println(bmSearch(badCharShiftMap, str, pat))
}