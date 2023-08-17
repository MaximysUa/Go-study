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