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

func BenchmarkRun10(b *testing.B) {
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i < 10; i++ {
		buf.Write([]byte(fmt.Sprintf("%d\n", rand.Intn(1000000))))
	}

	for n := 0; n < b.N; n++ {
		err := run(buf, ioutil.Discard)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkRun100(b *testing.B) {
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i < 100; i++ {
		buf.Write([]byte(fmt.Sprintf("%d\n", rand.Intn(1000000))))
	}

	for n := 0; n < b.N; n++ {
		err := run(buf, ioutil.Discard)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkRun1000(b *testing.B) {
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i < 100; i++ {
		buf.Write([]byte(fmt.Sprintf("%d\n", rand.Intn(1000000))))
	}

	for n := 0; n < b.N; n++ {
		err := run(buf, ioutil.Discard)
		if err != nil {
			b.Error(err)
		}
	}
}
