package main

import (
	"fmt"
	"strings"
	"testing"
	"testing/quick"
)

func TestOddMultipleOfThree(t *testing.T) {
	f := func(x int) bool {
		y := OddMultipleOfThree(x)
		return y%2 == 1 && y%3 == 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestStringConcat(t *testing.T) {
	f := func(s1 string, s2 string) bool {
		s := StringConcat(s1, s2)
		return s == s1+s2
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBazBar(t *testing.T) {
	f := func(foo *Foo) bool {

		checkBar := func(s string, bar string) bool {
			if bar == "" {
				return strings.Index(s, bar) == 0
			}
			return strings.Index(s, bar) > 0
		}

		checkLen := func(s string, substr1 string, substr2 string) bool {
			return len(s) == len(substr1)+len(substr2)
		}

		bb := BazBar(foo)
		baz := fmt.Sprintf("%d", foo.Baz)
		return strings.Index(bb, baz) == 0 && checkBar(bb, foo.Bar) && checkLen(bb, baz, foo.Bar)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
