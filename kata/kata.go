package kata

import (
	"strconv"
	"strings"
)

func FakeBin(x string) string {
	result := ""
	arr := strings.Split(x, "")

	for _, i := range arr {
		dig, _ := strconv.ParseInt(i, 10, 32)

		if dig < 5 {
			result += "0"
		} else {
			result += "1"
		}
	}
	return result
}
func NotMineFakeBin(x string) (result string) {
	for _, char := range x {
		if char < '5' {
			result += "0"
		} else {
			result += "1"
		}
	}
	return
}
func Hero(bullets, dragons int) bool {
	return dragons*2 <= bullets

}
func IsPalindrome(str string) bool {

	var reverseStr string
	runes := []rune(str)
	for i := len(runes) - 1; i >= 0; i-- {
		reverseStr += string(runes[i])
	}
	return strings.EqualFold(str, reverseStr)

}
func StringToArray(str string) []string {
	return strings.Split(str, " ")
}
func SetAlarm(employed, vacation bool) bool {
	return employed && !vacation
}
func Between(a, b int) []int {
	var result []int

	for i := a; i <= b; i++ {
		result = append(result, i)
	}
	return result
}
func CountSheeps(numbers []bool) int {
	var res int
	for _, i := range numbers {
		if i {
			res++
		}
	}
	return res
}
