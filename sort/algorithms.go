package sort

var comparisons, accesses int

// Bubble Sort Algorithm
func Bubble(arr []int, i int) ([]int, int, int) {
	if i != len(arr)-1 {
		if arr[i] > arr[i+1] {
			comparisons += 2
			accesses += 2
			arr = Swap(arr, i, i+1)
		}
	}
	return arr, comparisons, accesses
}

// Selection Sort Algorithm
func Selection(arr []int, i int) ([]int, int, int) {
	if i != len(arr)-1 {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			comparisons++
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr = Swap(arr, i, minIndex)
		}
	}
	return arr, comparisons, accesses
}

func Merge(arr []int, mid int) ([]int, int, int) {
	left := arr[:mid]
	right := arr[mid:]
	result := make([]int, 0, len(left)+len(right))
	var i, j int
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			accesses++
			result = append(result, left[i])
			i++
		} else {
			accesses++
			result = append(result, right[j])
			j++
		}
		comparisons++
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	accesses += len(left) + len(right)
	return result, comparisons, accesses
}

func CheckSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
		comparisons += 2
		accesses += 2
	}
	return true
}

func Swap(arr []int, i0, i1 int) []int {
	arr[i0], arr[i1] = arr[i1], arr[i0]
	accesses += 4
	return arr
}

func Argmax(arr []int) int {
	m := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[m] {
			m = i
		}
	}
	return m
}

func Amax(arr []int) int {
	m := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > m {
			m = arr[i]
		}
	}
	return m
}

func Argmin(arr []int) int {
	m := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[m] {
			m = i
		}
	}
	return m
}

func Amin(arr []int) int {
	m := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < m {
			m = arr[i]
		}
	}
	return m
}
