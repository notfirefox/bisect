package main

import "fmt"

type function func(int) int

type interval struct {
	a int
	b int
}

func bisect(fn function, i interval) int {
	var mid int = (i.a + i.b) / 2
	var fmid int = fn(mid)
	if fmid == 0 {
		return mid
	} else {
		var fa int = fn(i.a)
		var fb int = fn(i.b)
		if fa < 0 && fmid > 0 {
			return bisect(fn, interval{i.a, mid})
		}
		if fa > 0 && fmid < 0 {
			return bisect(fn, interval{i.a, mid})
		}
		if fb < 0 && fmid > 0 {
			return bisect(fn, interval{mid, i.b})
		}
		if fb > 0 && fmid < 0 {
			return bisect(fn, interval{mid, i.b})
		}
		return 666
	}
}

func main() {
	var i interval = interval{-1, 1}
	var z = bisect(func(x int) int { return (x * x) + (2 * x) }, i)
	fmt.Println(z)
}
