package parser_tests

import (
	"testing"

	"github.com/vincer2040/lexi-go/internal/parser"
	lexigo "github.com/vincer2040/lexi-go/pkg/lexi-go"
)

func TestParseStrings(t *testing.T) {
	input := []byte("$3\r\nfoo\r\n")
	p := parser.New(input, len(input))
	data, err := p.Parse()
	if err != nil {
		t.Fatalf("failed to parse %s\n", err)
	}
	if data.Type != lexigo.String {
		t.Fatalf("expected string, got %d\n", data.Type)
	}
	s := data.Data.(lexigo.LexiString)
	if s.Str != "foo" {
		t.Fatalf("expected string to be foo, got %s\n", s.Str)
	}
}

func TestParseIntegers(t *testing.T) {
	input := []byte(":1337\r\n")
	p := parser.New(input, len(input))
	data, err := p.Parse()
	if err != nil {
		t.Fatalf("failed to parse %s\n", err)
	}
	if data.Type != lexigo.Int {
		t.Fatalf("expected int, got %d\n", data.Type)
	}
	i := data.Data.(lexigo.LexiInt)
	if i.Integer != 1337 {
		t.Fatalf("expected int to be 1337, got %d\n", i.Integer)
	}
}

func TestParseDouble(t *testing.T) {
	input := []byte(",1337.1337\r\n")
	p := parser.New(input, len(input))
	data, err := p.Parse()
	if err != nil {
		t.Fatalf("failed to parse %s\n", err)
	}
	if data.Type != lexigo.Double {
		t.Fatalf("expected double, got %d\n", data.Type)
	}
	d := data.Data.(lexigo.LexiDouble)
	if d.Double != 1337.1337 {
		t.Fatalf("expected double to be 1337.1337, got %f\n", d.Double)
	}
}

func TestParseArray(t *testing.T) {
	input := []byte("*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n")
	p := parser.New(input, len(input))
	data, err := p.Parse()
	if err != nil {
		t.Fatalf("failed to parse %s\n", err)
	}
	if data.Type != lexigo.Array {
		t.Fatalf("expected array, got %d\n", data.Type)
	}
	exp := []string{"foo", "bar"}
	arr := data.Data.(lexigo.LexiArray)
	if len(arr.Array) != 2 {
		t.Fatalf("expected length to be 2, got %d\n", len(arr.Array))
	}

	for i, e := range exp {
		got := arr.Array[i]
		if got.Type != lexigo.String {
			t.Fatalf("expected type to be string, got %d\n", got.Type)
		}
		s := got.Data.(lexigo.LexiString)
		if s.Str != e {
			t.Fatalf("expected %s, got %s\n", e, s.Str)
		}
	}
}
