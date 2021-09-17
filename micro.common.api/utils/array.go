package utils

func UniqueArray(iArray []int64) []int64 {
	uniqueMap := make(map[int64]struct{})
	result := make([]int64, 0, len(iArray))

	for _, i := range iArray {
		if _, ok := uniqueMap[i]; !ok {
			result = append(result, i)
			uniqueMap[i] = struct{}{}
		}
	}

	return result
}

func IsInArray(item int, array []int) bool {
	for _, arrayItem := range array {
		if arrayItem == item {
			return true
		}
	}

	return false
}
