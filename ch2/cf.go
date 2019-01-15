package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tempconv/tempconv"
)

func comma(s string) string {
	var buf bytes.Buffer
	var sl = []byte(s)
	for i, ch := range sl {
		if len(sl) <= 3 {
			return string(sl)
		} else {
			buf.WriteByte(ch)
			if (i+1)%3 == 0 {
				buf.WriteString(",")
			}
		}
	}
	return buf.String()
}

func isAnagram(s1 string, s2 string) bool {
	var n1 = len(s1)
	var count1 = 0
	var n2 = len(s2)
	var count2 = 0
	if n1 != n2 {
		return false
	}
	if strings.Contains(s1, s2) {
		return true
	}
	for _, v := range []byte(s1) {
		if strings.Index(s2, string(v)) != -1 {
			count1++
		}
	}
	for _, v := range []byte(s2) {
		if strings.Index(s1, string(v)) != -1 {
			count2++
		}
	}
	return count1 == count2
}

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			os.Exit(1)
		}
		f := tempconv.Fahreneit(t)
		c := tempconv.Celsius(f)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FtoC(f), c, tempconv.CtoF(c))
	}

	fmt.Println(comma("12345"))
	fmt.Println(isAnagram("pippo", "poppi"))
}
