package popcount

import (
	"fmt"
	"testing"
)

func TestPopCount(t *testing.T) {
	fmt.Printf("%v\n", PopCount(10123408080))
}


func TestPopCountShift(t *testing.T) {
	fmt.Printf("%v\n", PopCountShift(10123408080))
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(10123408080)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(10123408080)
	}
}
