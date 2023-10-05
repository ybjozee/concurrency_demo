package main

import "testing"

func Benchmark_generateReportSynchronously(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateReportSynchronously()
	}
}

func Benchmark_generateReportConcurrently(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateReportConcurrently()
	}
}

//2 761 757 843 ns/op
//2 675 269 401 ns/op
//2 760 968 163 ns/op
//2 616 861 426 ns/op
//2 546 182 800 ns/op
//2 963 897 550 ns/op
//3 176 065 958 ns/op
//2 615 936 787 ns/op
//3 024 045 892 ns/op
//2 734 926 000 ns/op
