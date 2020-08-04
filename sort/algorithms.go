package sort

var (
	comparisons = 0
	accesses    = 0
)

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

/*
// Selection Sort Algorithm
func Selection(arr []int) int {
	return 0
}

// Merge Sort Algorithm
func Merge(arr []int) []int {
	if len(arr) > 1 {
		middle := int(math.Floor(float64(len(arr) / 2)))
		left := arr[:middle]
		right := arr[middle:]

		merge

	}
}
*/

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
