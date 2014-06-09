package itertools

import "testing"

func BenchmarkFilter(b *testing.B) {
	pred := func(i interface{}) bool {
		return i.(uint64)%2 == 1
	}

	for i := 0; i < b.N; i++ {
		Filter(pred, Uint64(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	}
}

func BenchmarkNoFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		results := make([]uint64, 0)

		for _, n := range input {
			if n%2 == 1 {
				results = append(results, n)
			}
		}
	}
}

func BenchmarkMap(b *testing.B) {
	mapper := func(i interface{}) interface{} {
		return len(i.(string))
	}

	for i := 0; i < b.N; i++ {
		Map(mapper, New("a", "ab", "abc", "abcd"))
	}
}

func BenchmarkNoMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"a", "ab", "abc", "abcd"}
		results := make([]int, 0)

		for _, w := range input {
			results = append(results, len(w))
		}
	}
}

func BenchmarkReduce(b *testing.B) {
	summer := func(memo interface{}, el interface{}) interface{} {
		return memo.(float64) + el.(float64)
	}

	for i := 0; i < b.N; i++ {
		Reduce(Float64(.1, .2, .3, .22), summer, float64(0))
	}
}

func BenchmarkNoReduce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []float64{.1, .2, .3, .22}
		result := 0.0

		for _, n := range input {
			result += n
		}
	}
}
