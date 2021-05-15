package sort

import (
	"sort"
)

func BubbleSort(elements []int) {
	keepWorking := true

	if keepWorking {
		keepWorking = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] < elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}

func Sort(elements []int) {
	sort.Ints(elements)
}
