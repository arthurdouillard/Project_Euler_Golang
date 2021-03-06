/*
Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/

package main

import "fmt"

var MAX int = 4000000

func main() {
	var f1 int = 1
	var f2 int = 2
	var f3 int

	var sum int = 2

	for f3 <= MAX {
		if f3%2 == 0 {
			sum += f3
		}

		f3 = f2 + f1
		f1 = f2
		f2 = f3
	}

	fmt.Println("Sum = ", sum)
}
