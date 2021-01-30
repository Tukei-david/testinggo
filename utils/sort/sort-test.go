package sort

import "testing"

// For testing if elements sorting has gone sucessfully.
func TestBubbleSort(t *testing.T) {

	// Test
	elements := []int{9, 7, 6, 4, 8, 5, 2, 3}

	// Exuction
	BubbleSort(elements)

	// Validation
	if elements[0] != 9 {
		t.Error("First element should be 9")
	}
	if elements[len(elements)-1] != 2 {
		t.Error("Last element should be 2")
	}

}
