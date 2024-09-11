package main

import (
	"fmt"
	"time"
    "slices"
)

func Bubblesort(s []float64) {
    n := len(s)
	for {
		swapped := false
		for i := range n-1 {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
				swapped = true
			}
		}
		if !swapped {
			return
		}
        n = n - 1
	}
}

// Fuer folgenden Quicksort inspierierte ich mich bei der C-Losung fuer
// Quicksort auf RosettaCode.
func Quicksort(s []float64) {
	var i, j int

	if len(s) < 2 {
		return
	}

	pivot := s[len(s)/2]
	for i, j = 0, len(s)-1; ; i, j = i+1, j-1 {
		for s[i] < pivot {
			i++
		}
		for s[j] > pivot {
			j--
		}

		if i >= j {
			break
		}

		s[i], s[j] = s[j], s[i]
	}
	Quicksort(s[:i])
	Quicksort(s[i:])
}

// func Quicksort(s []float64) {
// 	if len(s) <= 1 {
// 		return
// 	}
// 	div := Partition(s)
// 	Quicksort(s[:div])
// 	Quicksort(s[div+1:])
// }

// func Partition(s []float64) int {
// 	i := 0
// 	k := len(s) - 1
// 	j := k - 1
// 	pivot := s[k]

// 	for i < j {
// 		for i < j && s[i] <= pivot {
// 			i++
// 		}
// 		for i < j && s[j] > pivot {
// 			j--
// 		}
// 		if s[i] > s[j] {
// 			s[i], s[j] = s[j], s[i]
// 		}
// 	}
// 	if s[i] > pivot {
// 		s[i], s[k] = s[k], s[i]
// 	} else {
// 		i = k
// 	}
// 	return i
// }

func main() {
	s := make([]float64, 10_000)

    fmt.Printf("Try it first with Bubblesort...\n")
	Generate(s)
	t0 := time.Now()
	Bubblesort(s)
    if err := Check(s); err != nil {
        fmt.Printf("ERROR: %v\n", err)
    }
	d := time.Since(t0)
	fmt.Printf("  initial sorting: %v\n", d)
	t0 = time.Now()
	Bubblesort(s)
	d = time.Since(t0)
	fmt.Printf("  sort again     : %v\n", d)

    // fmt.Printf("First try with code from Wikipedia (Pivot at the border)...\n")
	// Generate(s)
	// t0 = time.Now()
	// Quicksort(s)
	// d = time.Since(t0)
	// fmt.Printf("  initial sorting: %v\n", d)
	// t0 = time.Now()
	// Quicksort(s)
	// d = time.Since(t0)
	// fmt.Printf("  sort again     : %v\n", d)

    fmt.Printf("Quicksort with Pivot in the middle of the slice...\n")
	// Generate(s)
    for i := range s {
        s[i] = 4.0
    }
	t0 = time.Now()
	Quicksort(s)
	d = time.Since(t0)
	fmt.Printf("  initial sorting: %v\n", d)
	t0 = time.Now()
	Quicksort(s)
	d = time.Since(t0)
	fmt.Printf("  sort again     : %v\n", d)

    fmt.Printf("Compared with the internal sorting method of Go...\n")
	// Generate(s)
    for i := range s {
        s[i] = 4.0
    }
	t0 = time.Now()
	slices.Sort(s)
	d = time.Since(t0)
	fmt.Printf("  initial sorting: %v\n", d)
	t0 = time.Now()
	slices.Sort(s)
	d = time.Since(t0)
	fmt.Printf("  sort again     : %v\n", d)


}
