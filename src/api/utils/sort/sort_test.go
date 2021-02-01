package sort

import (
	"fmt"
	"testing"
)

// TestBubbleSort will be used for testing if elements sorting has gone sucessfully.
func TestBubbleSortOrderDESC(t *testing.T) {

	// Init
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	fmt.Println(elements)
	fmt.Println(elements[len(elements)-1])

	// Exuction
	BubbleSort(elements)

	// Validation
	if elements[0] != 9 {
		t.Error("First element should be 9")
	}
	if elements[len(elements)-1] != 0 {
		t.Error("Last element should be 0")
	}

	fmt.Println(elements)

}
