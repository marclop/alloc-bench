## Nothing to see here. move along

    make bench
    goos: darwin
    goarch: amd64
    pkg: github.com/marclop/alloc-bench
    BenchmarkPlain-8                   	300000000	         5.88 ns/op	       0 B/op	       0 allocs/op
    BenchmarkPointer-8                 	2000000000	         0.56 ns/op	       0 B/op	       0 allocs/op
    BenchmarkPlainAlloc-8              	100000000	        10.2 ns/op	       0 B/op	       0 allocs/op
    BenchmarkPointerAlloc-8            	300000000	         4.48 ns/op	       0 B/op	       0 allocs/op
    BenchmarkPointerAllocAlternate-8   	300000000	         4.51 ns/op	       0 B/op	       0 allocs/op
    PASS
