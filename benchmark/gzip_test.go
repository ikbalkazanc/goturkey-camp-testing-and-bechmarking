package main

import "testing"

func BenchmarkGzipWithoutPool(b *testing.B) {
	gzip := NewGzipWithoutPool()

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := gzip.Zip("Go Turkiye Egitim Kampi - 203 Test & Benchmarks")

		if err != nil {
			b.Fatal(err)
		}

	}
}

func BenchmarkGzipWithPool(b *testing.B) {
	gzip := NewGzipWithPool()

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := gzip.Zip("Go Turkiye Egitim Kampi - 203 Test & Benchmarks")

		if err != nil {
			b.Fatal(err)
		}

	}
}
