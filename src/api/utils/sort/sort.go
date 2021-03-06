package sort

// BubbleSort will be For sorting out the elements
func BubbleSort(elements []int) {
	keepWorking := true

	for keepWorking {
		keepWorking = false

		for i := 0; i < len(elements); i++ {
			if elements[i] < elements[i+1] {
				keepWorking = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}
