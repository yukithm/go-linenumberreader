package linenumberreader_test

import (
	"io/ioutil"
	"strings"
	"testing"

	linenumberreader "github.com/yukithm/go-linenumberreader"
)

func TestEndWithLF(t *testing.T) {
	src := strings.NewReader("foo\nbar\nbaz\n")
	r := linenumberreader.NewLineNumberReader(src)
	_, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	got := r.LineNumber
	want := 2
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestEndWithCR(t *testing.T) {
	src := strings.NewReader("foo\rbar\rbaz\r")
	r := linenumberreader.NewLineNumberReader(src)
	_, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	got := r.LineNumber
	want := 2
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestEndWithCRLF(t *testing.T) {
	src := strings.NewReader("foo\r\nbar\r\nbaz\r\n")
	r := linenumberreader.NewLineNumberReader(src)
	_, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	got := r.LineNumber
	want := 2
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestEndWithoutBreak(t *testing.T) {
	src := strings.NewReader("foo\nbar\r\nbaz")
	r := linenumberreader.NewLineNumberReader(src)
	_, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	got := r.LineNumber
	want := 2
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNoBreak(t *testing.T) {
	src := strings.NewReader("foo")
	r := linenumberreader.NewLineNumberReader(src)
	_, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	got := r.LineNumber
	want := 0
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}
