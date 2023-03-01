package algo

// MergeSort принимает массив целых чисел в качестве
// входных данных и делит его на две половины. Переменные `left`
// и `right` хранят две половины, которые затем рекурсивно
// сортируются перед объединением обратно в отсортированном порядке.
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return Merge(left, right)
}

// Merge объединяет две половинки вместе в отсортированном порядке.
// Он использует два индекса для отслеживания левой и правой сторон
// и добавляет элементы с двух сторон соответственно.
// Наконец, возвращается результирующий отсортированный массив.
func Merge(left, right []int) []int {
	ret := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(ret, right...)
		}
		if len(right) == 0 {
			return append(ret, left...)
		}
		if left[0] <= right[0] {
			ret = append(ret, left[0])
			left = left[1:]
		} else {
			ret = append(ret, right[0])
			right = right[1:]
		}
	}

	return ret
}

// InsertionSort - функция, которая сортирует массив с использованием алгоритма сортировки вставками
func InsertionSort(items []int) []int {
	// получаем длину массива
	size := len(items)

	// организуем цикл по индексу i
	for i := 1; i < size; i++ {
		// устанавливаем индекс j равным i и идем по массиву с конца
		j := i
		key := items[j]
		for j > 0 && items[j-1] > key {
			// если условие выполняется, то обмениваем значения
			items[j] = items[j-1]
			// уменьшаем индекс j
			j--
		}
		items[j] = key
	}

	return items
}
