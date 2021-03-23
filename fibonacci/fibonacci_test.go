package fibonacchi

import "testing"

func TestFibonacciR(t *testing.T) {
	testCase := []struct {
		n    uint8
		want uint64
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {10, 55}, {42, 267914296},
	}

	for _, test := range testCase {
		if result := FibonacciR(test.n); result != test.want {
			t.Errorf("Invalid Fibonacci value for N: %d, got: %d, want: %d", test.n, result, test.want)
		}
	}
}

func TestFibonacciM(t *testing.T) {
	testCase := []struct {
		n    uint8
		want uint64
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {10, 55}, {42, 267914296},
	}

	for _, test := range testCase {
		if result := FibonacciM(test.n); result != test.want {
			t.Errorf("Invalid Fibonacci value for N: %d, got: %d, want: %d", test.n, result, test.want)
		}
	}
}

func BenchmarkFibonacciR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciR(30)
	}
}

func BenchmarkFibonacciM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciM(30)
	}
}
