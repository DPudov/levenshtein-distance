package tests

import (
	"levenshtein-distance/levenshtein"
	"levenshtein-distance/util"
	"testing"
)

func benchmarkLevenshteinIterative(b *testing.B, count int) {
	b.StopTimer()
	first := util.RandomString(count)
	second := util.RandomString(count)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		levenshtein.LevenshteinIterative(first, second)
	}
}

func benchmarkLevenshteinRecursive(b *testing.B, count int) {
	b.StopTimer()
	first := util.RandomString(count)
	second := util.RandomString(count)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		levenshtein.LevenshteinRecursiveOptimized(first, second)
	}
}

func benchmarkLevenshteinDamerau(b *testing.B, count int) {
	b.StopTimer()
	first := util.RandomString(count)
	second := util.RandomString(count)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		levenshtein.LevenshteinDamerau(first, second)
	}
}

func BenchmarkLevenshteinIterative100(b *testing.B) {
	benchmarkLevenshteinIterative(b, 100)
}

func BenchmarkLevenshteinIterative200(b *testing.B) {
	benchmarkLevenshteinIterative(b, 200)
}

func BenchmarkLevenshteinIterative300(b *testing.B) {
	benchmarkLevenshteinIterative(b, 300)
}

func BenchmarkLevenshteinIterative400(b *testing.B) {
	benchmarkLevenshteinIterative(b, 400)
}

func BenchmarkLevenshteinIterative500(b *testing.B) {
	benchmarkLevenshteinIterative(b, 500)
}

func BenchmarkLevenshteinIterative600(b *testing.B) {
	benchmarkLevenshteinIterative(b, 600)
}

func BenchmarkLevenshteinIterative700(b *testing.B) {
	benchmarkLevenshteinIterative(b, 700)
}

func BenchmarkLevenshteinIterative800(b *testing.B) {
	benchmarkLevenshteinIterative(b, 800)
}

func BenchmarkLevenshteinIterative900(b *testing.B) {
	benchmarkLevenshteinIterative(b, 900)
}

func BenchmarkLevenshteinIterative1000(b *testing.B) {
	benchmarkLevenshteinIterative(b, 1000)
}

func BenchmarkLevenshteinRecursive100(b *testing.B) {
	benchmarkLevenshteinRecursive(b, 100)
}

func BenchmarkLevenshteinRecursive200(b *testing.B) {
	benchmarkLevenshteinRecursive(b, 200)
}

func BenchmarkLevenshteinRecursive300(b *testing.B) {
	benchmarkLevenshteinRecursive(b, 300)
}

func BenchmarkLevenshteinRecursive400(b *testing.B) {
	benchmarkLevenshteinRecursive(b, 400)
}

func BenchmarkLevenshteinRecursive500(b *testing.B) {
	benchmarkLevenshteinRecursive(b, 500)
}

func BenchmarkLevenshteinRecursive600(b *testing.B) {
	benchmarkLevenshteinRecursive(b, 600)
}

func BenchmarkLevenshteinRecursive700(b *testing.B) {
	benchmarkLevenshteinRecursive(b, 700)
}

func BenchmarkLevenshteinRecursive800(b *testing.B) {
	benchmarkLevenshteinRecursive(b, 800)
}

func BenchmarkLevenshteinRecursive900(b *testing.B) {
	benchmarkLevenshteinRecursive(b, 900)
}

func BenchmarkLevenshteinRecursive1000(b *testing.B) {
	benchmarkLevenshteinRecursive(b, 1000)
}

func BenchmarkLevenshteinDamerau100(b *testing.B) {
	benchmarkLevenshteinDamerau(b, 100)
}

func BenchmarkLevenshteinDamerau200(b *testing.B) {
	benchmarkLevenshteinDamerau(b, 200)
}

func BenchmarkLevenshteinDamerau300(b *testing.B) {
	benchmarkLevenshteinDamerau(b, 300)
}

func BenchmarkLevenshteinDamerau400(b *testing.B) {
	benchmarkLevenshteinDamerau(b, 400)
}

func BenchmarkLevenshteinDamerau500(b *testing.B) {
	benchmarkLevenshteinDamerau(b, 500)
}

func BenchmarkLevenshteinDamerau600(b *testing.B) {
	benchmarkLevenshteinDamerau(b, 600)
}

func BenchmarkLevenshteinDamerau700(b *testing.B) {
	benchmarkLevenshteinDamerau(b, 700)
}

func BenchmarkLevenshteinDamerau800(b *testing.B) {
	benchmarkLevenshteinDamerau(b, 800)
}

func BenchmarkLevenshteinDamerau900(b *testing.B) {
	benchmarkLevenshteinDamerau(b, 900)
}

func BenchmarkLevenshteinDamerau1000(b *testing.B) {
	benchmarkLevenshteinDamerau(b, 1000)
}

//func TestLevenshteinRecursive(t *testing.T) {
//	defer measuring.ElapsedTime(time.Now(), "test")
//	first := util.RandomString(50)
//	second := util.RandomString(50)
//	//start := time.Now()
//	//r1 := levenshtein.LevenshteinRecursive(first, second)
//	//since := time.Since(start)
//	start2 := time.Now()
//	_ = levenshtein.LevenshteinRecursiveOptimized(first, second)
//	since2 := time.Since(start2)
//	//if r1 != r2 {
//	//	t.Error(r1, r2)
//	//}
//	log.Println("PASSED ",  since2)
//}
