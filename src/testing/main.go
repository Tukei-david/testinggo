package main

import "fmt"

func checkSize(a int) string {
	switch {
	case a < 0:
		return "Negative"
	case a == 0:
		return "Zero"
	case a < 10:
		return "small"
	case a < 100:
		return "big"
	case a < 1000:
		return "huge"
	}
	return "Enormus"
}

func main() {
	fmt.Printf(checkSize(10))
}
