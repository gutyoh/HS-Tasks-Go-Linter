// https://stepik.org/edit-lesson/763847/step/8

// Linter log:
/*
NONE -- Passed! ✅
*/

package main

import "fmt"

func main() {
	var base, inc int
	fmt.Scan(&base, &inc)

	increment := newIncrement(base)
	fmt.Println(increment(inc))
}

func newIncrement(base int) func(int) int {
	number := base
	return func(inc int) int {
		number += inc
		return number
	}
}
