package test

import "testing"

type Struct struct {
	s1, s2, s3, s4, s5, s6 string
	i1, i2, i3, i4, i5, i6 int
}

const SIZE = 1024

func BenchmarkSliceOfPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]*Struct, 0, SIZE)
		for j := 0; j < SIZE; j++ {
			_ = &Struct{}
			slice = append(slice, &Struct{})
		}
	}
}

func BenchmarkSliceOfStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]Struct, 0, SIZE)
		for j := 0; j < SIZE; j++ {
			slice = append(slice, Struct{})
		}
	}
}
