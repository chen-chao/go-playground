package builtin

import (
	"fmt"
	"testing"
)

func TestStringConvertion(t *testing.T) {
	s := "abcAbc123"
	// keep the binary representation unchanged like C when converting data types?
	// char to uint64
	fmt.Println("char to uint64")
	for _, c := range s {
		fmt.Printf("%c -> %d\n", c, uint64(c))
	}

	// char to float64
	fmt.Println("char to float64")
	for _, c := range s {
		fmt.Printf("%c -> %.1f\n", c, float64(c))
	}

}
