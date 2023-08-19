package kata



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
			if arr[i] == -arr[j]{
				flag = true
			}
		}
		if !flag {
			return arr[i]
		}
	}
	return 0
}

//TODO - разобрать эту хуйню
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