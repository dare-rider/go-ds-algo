package main

import (
	"bytes"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var str string
	fmt.Scanf("%s", &str)
	fmt.Print(formatToCurrenct(str))
}

func formatToCurrenct(str string) string{
	var res []string
	strTof, _ := strconv.ParseFloat(str, 64)
	str = fmt.Sprintf("%.2f", strTof)
	strs := strings.Split(str, ".")
	num := strs[0]
	dec := strs[1]
	i := len(num) - 1
	for i >= 0{
		j := i - 2
		if j < 0 {
			j = 0
		}
		res = append(res, num[j:i+1])
		i -= 3
	}
	var reStr bytes.Buffer
	i = len(res) - 1
	for i >= 0 {
		reStr.WriteString(res[i])
		if i != 0 {
			reStr.WriteString(",")
		}
		i -= 1
	}
	if dec != "" {
		reStr.WriteString(fmt.Sprintf(".%s", dec))
	}
	return reStr.String()
}
