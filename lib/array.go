package lib

func ArrayMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func ArrayMin(arr []int) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func ArraySum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func ArrayLast[K comparable](arr []K) K {
	return arr[len(arr)-1]
}

func ArrayRemoveItem[K comparable](arr []K, value K) []K {
	index := ArrayIndexOf(arr, value)
	if index == -1 {
		return arr
	} else if index == len(arr)-1 {
		return arr[:len(arr)-1]
	} else {
		return append(arr[:index], arr[index+1:]...)
	}
}

func ArrayContains[K comparable](arr []K, value K) bool {
	return ArrayIndexOf(arr, value) != -1
}

func ArrayIndexOf[K comparable](arr []K, value K) int {
	for i := range arr {
		if arr[i] == value {
			return i
		}
	}
	return -1
}
