package main

import "fmt"

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i, v := range x {
		if v != y[i] {
			return false
		}
	}

	return true
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func main() {
	x := []string{"a", "c"}
	y := []string{"a", "b"}
	fmt.Println(equal(x, y))

	z := []int{1, 2, 3, 4, 5}
	reverse(z)
	fmt.Println(z)

	var q, p []int
	for i := 0; i < 10; i++ {
		p = appendInt(q, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(p), p)
		q = p
	}

}
