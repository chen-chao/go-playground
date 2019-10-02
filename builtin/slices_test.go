package builtin

import (
	"fmt"
	"testing"
)

func TestSlices(t *testing.T) {
	// make can create fixed length slices
	s := make([]int, 5) // len(s) = 5, cap(s) = 5
	printSlice("s", s)
	// reference s with index
	fmt.Println("s[3] = ", s[3])
	// error: index out of range
	// fmt.Println("s[10] = ", s[10])

	// index reference will exclude elements before the first index
	b := s[:2] // len(b) = 2, cap(b) = 10
	printSlice("b", b)

	c := b[1:5] // len(c) = 4, cap(c) = 9
	printSlice("c", c)

	r := make([]int, 3)
	// append will expand slice by two times
	for i := 0; i < 10; i++ {
		r = append(r, i)
		printSlice("r", r)
	}

	// ordinary slice
	var p []int
	for i := 0; i < 10; i++ {
		p = append(p, i)
		printSlice("p", p)
	}

	// deletion
	printSlice("r", r)
	for i := 0; i < 10; i++ {
		r = r[0 : len(r)-1]
		printSlice("r", r)
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
