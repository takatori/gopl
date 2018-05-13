package popcount

import (
	"fmt"
	"testing"
)

func TestPopCount(t *testing.T) {
	fmt.Printf("%v\n", PopCount(10123408080))
}


func TestPopCountShift(t *testing.T) {
	fmt.Printf("%v\n", PopCountBitClear(10123408080))
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(10123408080)
	}
}

func BenchmarkPopCountBitClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountBitClear(10123408080)
	}
}
