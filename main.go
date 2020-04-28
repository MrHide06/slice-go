package main

import "fmt"

func main() {
	// Creates an array of size 10, slices it till index 5, and returns the slice reference
	s := make([]int, 1)
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	for angka := 1; angka < 10; angka++ {
		i := 0
		for bilangan := 1; bilangan < 10; bilangan++ {
			if angka%bilangan == 0 {
				i++
			}
		}
		if (i == 2) && (angka != 1) {
			s = append(s, angka)
			fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
		}
	}

}
