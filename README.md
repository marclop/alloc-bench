## Nothing to see here. move along

    make bench
    goos: darwin
    goarch: amd64
    pkg: github.com/marclop/alloc-bench
    BenchmarkPlain-8          	1000000000	         2.47 ns/op	       0 B/op	       0 allocs/op
    BenchmarkPointer-8        	2000000000	         0.54 ns/op	       0 B/op	       0 allocs/op
    BenchmarkPlainAlloc-8     	1000000000	         2.46 ns/op	       0 B/op	       0 allocs/op
    BenchmarkPointerAlloc-8   	500000000	         3.77 ns/op	       0 B/op	       0 allocs/op
    PASS
