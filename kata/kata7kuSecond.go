package kata

import (
	"fmt"
	"math"
	"strings"
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

/*
	func GetCount(str string) (count int) {
		vowels := []rune("aeiou")
		for _, i := range str {
			switch i {
			case vowels[0]:
				count++
			case vowels[1]:
				count++
			case vowels[2]:
				count++
			case vowels[3]:
				count++
			case vowels[4]:
				count++

			}

		}
		return count
	}
*/
func GetCount(str string) (count int) {
	for _, c := range str {
		switch c {
		case 'a', 'A', 'E', 'I', 'O', 'U', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}

//Прикольно можно удалять буквы из слова
/*func Disemvowel(comment string) string {
	for _,c := range "aiueoAIUEO"{
		comment = strings.ReplaceAll(comment, string(c), "")
	}
	return comment
}*/
func Disemvowel(comment string) string {
	var sb strings.Builder
	for _, c := range comment {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			continue
		default:
			sb.WriteRune(c)

		}
	}
	return sb.String()
}
