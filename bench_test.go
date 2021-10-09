package errors

import "testing"

func BenchmarkError1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New("an_error")
	}
}

func BenchmarkError10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New("an_error")
	}
}

func BenchmarkWrappedError1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Wrap(New("an_error"), "wrapped")
	}
}

func BenchmarkWrappedError10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Wrap(New("an_error"), "wrapped")
	}
}
