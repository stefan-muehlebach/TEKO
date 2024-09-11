package main

import (
	"slices"
	"math/rand"
	"testing"
)

const (
	sliceLen = 10_000
)

func SetRand(s []float64) {
	for i := range s {
		s[i] = rand.Float64()
	}
}

func BenchmarkSetRand(b *testing.B) {
	s := make([]float64, sliceLen)
	for range b.N {
		SetRand(s)
	}
}

func BenchmarkQuicksort(b *testing.B) {
	s := make([]float64, sliceLen)
	for range b.N {
		SetRand(s)
		Quicksort(s)
	}
}

func BenchmarkBubblesort(b *testing.B) {
	s := make([]float64, sliceLen)
	for range b.N {
		SetRand(s)
		Bubblesort(s)
	}
}

func BenchmarkInternalSort(b *testing.B) {
	s := make([]float64, sliceLen)
	for range b.N {
		SetRand(s)
		slices.Sort(s)
	}
}


