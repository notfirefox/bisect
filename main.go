package main

import "fmt"

type function func(int) int

type interval struct {
	a int
	b int
}

func function_impl(x int) int {
	return (x * x) + (2 * x)
}

func bisect(fn function, i interval) int {
	var mid int = (i.a + i.b) / 2
	var fmid int = fn(mid)
	if fmid == 0 {
		return mid
	} else if fmid > 0 {
		return bisect(fn, interval{i.a, mid})
	} else {
		return bisect(fn, interval{mid, i.b})
	}
}

func main() {
	var f function = function_impl
	var i interval = interval{-3, -1}
	fmt.Println(bisect(f, i))
}
