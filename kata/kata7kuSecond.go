package kata

import (
	"fmt"
	"math"
)

func IntRac(n uint64, guess int) int {
	oldGuess := uint64(guess)
	var counter int

	for {
		counter++
		x := (n/oldGuess + oldGuess) / 2
		fmt.Println(x)
		if x == oldGuess {
			return counter
		}

		oldGuess = x
	}
	return 0
}
func Solution(n, b int) []int {
	var res []int
	if b == 0 || n == 0 {
		return res
	}
	for i := 1; i <= int(math.Pow(2.0, (float64(n)))-1); i++ {
		x := fmt.Sprintf("%b", i)
		fmt.Println(x)
		if x[b] == 1 {
			res = append(res, i)
		}
	}
	return res
}
