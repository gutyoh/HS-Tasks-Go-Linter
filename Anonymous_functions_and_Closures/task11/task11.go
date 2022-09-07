// https://stepik.org/edit-lesson/763847/step/11

// Linter log:
/*
task11\task11.go:6:33: mnd: Magic number: 100, in <argument> detected (gomnd)
        constraint := newConstraint(0, 100)
*/

package main

import "fmt"

func main() {
	constraint := newConstraint(0, 100)

	var number int
	fmt.Scan(&number)

	fmt.Println(constraint(number))
}

func newConstraint(a, b int) func(num int) int {
	minValue := a
	maxValue := b

	return func(num int) int {
		if num < minValue {
			return minValue
		}
		if num > maxValue {
			return maxValue
		}
		return num
	}
}
