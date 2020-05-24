package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"testing"
)

func BenchmarkRun1(b *testing.B) {
	buf := bytes.NewBuffer([]byte{})
	buf.Write([]byte("123456\n"))

	for n := 0; n < b.N; n++ {
		err := run(buf, ioutil.Discard)
		if err != nil {
			b.Error(err)
		}
	}
}

func benchmarkRunN(b *testing.B, n int) {
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i < n; i++ {
		buf.Write([]byte(fmt.Sprintf("%d\n", rand.Intn(1000000))))
	}

	for n := 0; n < b.N; n++ {
		err := run(buf, ioutil.Discard)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkRun10(b *testing.B) {
	benchmarkRunN(b, 10)
}

func BenchmarkRun100(b *testing.B) {
	benchmarkRunN(b, 100)
}

func BenchmarkRun1k(b *testing.B) {
	benchmarkRunN(b, 1000)
}

func BenchmarkRun10k(b *testing.B) {
	benchmarkRunN(b, 10000)
}

func BenchmarkRun100k(b *testing.B) {
	benchmarkRunN(b, 100000)
}

func BenchmarkRun1m(b *testing.B) {
	benchmarkRunN(b, 1000000)
}

func BenchmarkRun2m(b *testing.B) {
	benchmarkRunN(b, 2000000)
}
