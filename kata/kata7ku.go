package kata

import (
	"fmt"
	"sort"
)

func Cats(start, finish int) int {
	var dig int
	for start < finish {
		if start+3 <= finish {
			start += 3
			dig++
		} else if start+1 <= finish {
			start++
			dig++

		}
	}
	return dig
}

func ArrayElements(arr []int) int {
	for i := 0; i < len(arr); i++ {
		flag := false
		for j := 0; j < len(arr); j++ {
			if arr[i] == -arr[j] {
				flag = true
			}
		}
		if !flag {
			return arr[i]
		}
	}
	return 0
}

// TODO - разобрать эту хуйню
func NotMineArrayElements(arr []int) int {
	hash := make(map[int]bool)
	ans := 0
	for _, entry := range arr {
		if _, value := hash[entry]; !value {
			hash[entry] = true
			ans += entry
		}
	}
	return ans
}
func RacePodium(blocks int) [3]int {
	switch {
	case blocks == 7: // the only exception to the formula
		return [3]int{2, 4, 1}
	case blocks%3 == 0:
		return [3]int{blocks / 3, blocks/3 + 1, blocks/3 - 1}
	case blocks%3 == 1:
		return [3]int{blocks/3 + 1, blocks/3 + 2, blocks/3 - 2}
	case blocks%3 == 2:
		return [3]int{blocks/3 + 1, blocks/3 + 2, blocks/3 - 1}
	}
	panic("invalid input")
}
func MaxMultiple(d, b int) int {
	for i := b; i > 0; i-- {
		if i%d == 0 {
			return i
		}
	}
	return 0
	//variant of solution
	// return (b / d) * d
}
func SortNumbers(numbers []int) []int {

	sort.Ints(numbers)
	return numbers
}

func PrinterError(s string) string {
	var num int
	for _, value := range s {
		if value > 'm' {
			num++
		}
	}
	return fmt.Sprintf("%d/%d", num, len(s))
}
func Pendulum(values []int) []int {
	sort.Ints(values)
	res := make([]int, len(values))
	mid := (len(values) - 1) / 2
	for i, val := range values {
		if i < len(values) {
			if i%2 == 0 {
				res[mid-i] = val
				mid -= i
			} else {
				res[mid+i] = val
				mid += i
			}
		}
	}
	return res
}
// its better solution then mine
//func Pendulum(values []int) (dst []int) {
//	sort.Ints(values)
//	for i, v := range values {
//		if i % 2 == 0 {
//			dst = append([]int{v}, dst...)
//		} else {
//			dst = append(dst, v)
//		}
//	 }
//	 
//	 return
//}

