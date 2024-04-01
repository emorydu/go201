package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var s = []string{
	"c",
	"go",
	"hello, world",
}

func concatStringByOperator(s []string) string {
	var st string
	for _, v := range s {
		st += v
	}

	return st
}

func concatStringByJoin(s []string) string {
	return strings.Join(s, "")
}

func concatStringByStringBuilder(s []string) string {
	var b strings.Builder
	for _, v := range s {
		b.WriteString(v)
	}

	return b.String()
}

func concatStringByBytesBuffer(s []string) string {
	var b bytes.Buffer
	for _, v := range s {
		b.WriteString(v)
	}

	return b.String()
}

func concatStringByBytesBufferWithInitSize(s []string) string {
	buf := make([]byte, 0, 64)
	b := bytes.NewBuffer(buf)
	for _, v := range s {
		b.WriteString(v)
	}

	return b.String()
}

func concatStringBySprintf(s []string) string {
	var st string
	for _, v := range s {
		fmt.Sprintf("%s%s", s, v)
	}

	return st
}

func BenchmarkConcatStringByOperator(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByOperator(s)
	}
}

func BenchmarkConcatStringBySprintf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringBySprintf(s)
	}
}

func BenchmarkConcatStringByStringsBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByStringBuilder(s)
	}
}

func BenchmarkConcatStringByJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByJoin(s)
	}
}

func BenchmarkConcatStringByBytesBufferWithInitSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByBytesBufferWithInitSize(s)
	}
}
