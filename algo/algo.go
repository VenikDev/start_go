package algo

// FindMaximum ищет максимальное число в слайсе.
// Возвращается максимальное число.
func FindMaximum(numbers []int) int {
	max := numbers[0]

	for _, num := range numbers {
		if num > max {
			max = num
		}
	}

	return max
}

// Add вычисляет сумму x и y и возвращает сумму.
func Add(x, y int64) int64 {
	return x + y
}

// TwoLargest ищет два наибольших числа в слайсе целых чисел.
// Возвращается два максимальных числа.
// Если массив пустой - возвращается (-1, -1).
// Если в массиве один элемент – второй результат будет равен -1.
func TwoLargest(arr []int) (int, int) {
	var largest, secondLargest int

	if len(arr) == 0 {
		return -1, -1
	}

	if len(arr) == 1 {
		return arr[0], arr[0]
	}

	largest, secondLargest = arr[0], arr[1]

	if largest < secondLargest {
		largest, secondLargest = secondLargest, largest
	}

	for i := 2; i < len(arr); i++ {
		if arr[i] > largest {
			secondLargest = largest
			largest = arr[i]
		} else if arr[i] > secondLargest && arr[i] != largest {
			secondLargest = arr[i]
		}
	}

	return largest, secondLargest
}

// MinMax ищет минимальное и максимальное значение в срезе.
// Возвращает минимальное и максимальное значения.
func MinMax(arr []int) (int, int) {
	var min, max int

	if len(arr) == 0 {
		return -1, -1
	}

	if len(arr) == 1 {
		return arr[0], arr[0]
	}

	min, max = arr[0], arr[1]

	for i := 2; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		} else if arr[i] > max {
			max = arr[i]
		}
	}

	return min, max
}

func Sum(arr []int) int {
	// define a variable to store the sum
	var sum int

	// loop over each element of the array
	for _, value := range arr {
		// add value to the sum
		sum += value
	}

	// return the sum
	return sum
}

// Reverse takes an array of integers and reverses its order
func Reverse(list []int) {
	//i is the starting index and j is the ending index of the list
	for i, j := 0, len(list)-1; i < j; {
		//swaps the values at the indexes i and j
		list[i], list[j] = list[j], list[i]
		//increment i and decrement j
		i++
		j--
	}
}
