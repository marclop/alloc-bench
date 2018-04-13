package alloctest

import (
	"testing"
)

var p string
var pp string
var pa string
var ppa string
var ppaa string

func BenchmarkPlain(b *testing.B) {
	var r string
	t := T{"1123123123", "212312312", "3", 0}
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		t.I = n
		r = plain(t)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	p = r
}

func BenchmarkPointer(b *testing.B) {
	var r string
	t := &T{"1123123123", "212312312", "3", 0}
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		t.I = n
		r = pointer(t)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	pp = r
}

func BenchmarkPlainAlloc(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = plain(T{"1123123123", "212312312", "3", n})
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	pa = r
}

func BenchmarkPointerAlloc(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = pointer(&T{"1123123123", "212312312", "3", n})
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	ppa = r
}

func BenchmarkPointerAllocAlternate(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		t := T{"1123123123", "212312312", "3", n}
		r = pointer(&t)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	ppaa = r
}
