// https://stepik.org/edit-lesson/763847/step/9

// Linter log:
/*
task9\task9.go:10:3: assignOp: replace `number = number * 2` with `number *= 2` (gocritic)
                number = number * 2
                ^
*/

package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	increment := func() {
		number = number * 2
		number++
	}

	increment()
	fmt.Println(number)
}
