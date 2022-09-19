package main

import (
	"testing"
)

var testArr2 = []int{8, 2, 3, 4, 5, 6, 7, 9, 12, 13, 14, 15}

var tstArr = []int{10, 9, 2, 8, 3, 7, 4, 6, 5, 0, 22, -22}

func BenchmarkSimplest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		choseSort(testArr2)
	}
}

func BenchmarkSimplest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fastSort(testArr2)
	}
}
