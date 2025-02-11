// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmdline

import (
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/pflag"
)

func setUpSSFlagSet(ssp *stringSliceValue) *pflag.FlagSet {
	f := pflag.NewFlagSet("test", pflag.ContinueOnError)
	f.Var(ssp, "ss", "Command separated list!")
	return f
}

func setUpSSFlagSetWithDefault(ssp *stringSliceValue) *pflag.FlagSet {
	f := pflag.NewFlagSet("test", pflag.ContinueOnError)
	ssp.Replace([]string{"default", "values"})
	f.Var(ssp, "ss", "Command separated list!")
	return f
}

func TestEmptySS(t *testing.T) {
	var ss stringSliceValue
	f := setUpSSFlagSet(&ss)
	err := f.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getSS := ss.GetSlice()
	if len(getSS) != 0 {
		t.Fatalf("got ss %v with len=%d but expected length=0", getSS, len(getSS))
	}
}

func TestEmptySSValue(t *testing.T) {
	var ss stringSliceValue
	f := setUpSSFlagSet(&ss)
	err := f.Parse([]string{"--ss="})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getSS := ss.GetSlice()
	if len(getSS) != 0 {
		t.Fatalf("got ss %v with len=%d but expected length=0", getSS, len(getSS))
	}
}

func TestSS(t *testing.T) {
	var ss stringSliceValue
	f := setUpSSFlagSet(&ss)

	vals := []string{"one", "two", "4", "3"}
	arg := fmt.Sprintf("--ss=%s", strings.Join(vals, ","))
	err := f.Parse([]string{arg})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getSS := ss.GetSlice()
	for i, v := range getSS {
		if vals[i] != v {
			t.Fatalf("expected ss[%d] to be %s from GetStringSlice but got: %s", i, vals[i], v)
		}
	}
}

func TestSSDoubleQuotes(t *testing.T) {
	var ss stringSliceValue
	f := setUpSSFlagSet(&ss)

	vals := []string{`a="one"`, `"two"=b`, `three`}
	arg := fmt.Sprintf("--ss=%s", strings.Join(vals, ","))
	err := f.Parse([]string{arg})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getSS := ss.GetSlice()
	for i, v := range getSS {
		if vals[i] != v {
			t.Fatalf("expected ss[%d] to be %s from GetStringSlice but got: %s", i, vals[i], v)
		}
	}
}

func TestSSDefault(t *testing.T) {
	var ss stringSliceValue
	f := setUpSSFlagSetWithDefault(&ss)

	vals := []string{"default", "values"}

	err := f.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getSS := ss.GetSlice()
	for i, v := range getSS {
		if vals[i] != v {
			t.Fatalf("expected ss[%d] to be %s from GetStringSlice but got: %s", i, vals[i], v)
		}
	}
}

func TestSSWithDefault(t *testing.T) {
	var ss stringSliceValue
	f := setUpSSFlagSetWithDefault(&ss)

	vals := []string{"one", "two", "4", "3"}
	arg := fmt.Sprintf("--ss=%s", strings.Join(vals, ","))
	err := f.Parse([]string{arg})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getSS := ss.GetSlice()
	for i, v := range getSS {
		if vals[i] != v {
			t.Fatalf("expected ss[%d] to be %s from GetStringSlice but got: %s", i, vals[i], v)
		}
	}
}

func TestSSCalledTwice(t *testing.T) {
	var ss stringSliceValue
	f := setUpSSFlagSet(&ss)

	in := []string{"one,two", "three"}
	expected := []string{"one", "two", "three"}
	argfmt := "--ss=%s"
	arg1 := fmt.Sprintf(argfmt, in[0])
	arg2 := fmt.Sprintf(argfmt, in[1])
	err := f.Parse([]string{arg1, arg2})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	values := ss.GetSlice()

	if len(expected) != len(values) {
		t.Fatalf("expected number of values to be %d but got: %d", len(expected), len(values))
	}
	for i, v := range values {
		if expected[i] != v {
			t.Fatalf("expected got ss[%d] to be %s but got: %s", i, expected[i], v)
		}
	}
}

func TestSSWithComma(t *testing.T) {
	var ss stringSliceValue
	f := setUpSSFlagSet(&ss)

	in := []string{`"one,two"`, `"three"`, `"four,five",six`}
	expected := []string{"one,two", "three", "four,five", "six"}
	argfmt := "--ss=%s"
	arg1 := fmt.Sprintf(argfmt, in[0])
	arg2 := fmt.Sprintf(argfmt, in[1])
	arg3 := fmt.Sprintf(argfmt, in[2])
	err := f.Parse([]string{arg1, arg2, arg3})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	values := ss.GetSlice()

	if len(expected) != len(values) {
		t.Fatalf("expected number of values to be %d but got: %d", len(expected), len(values))
	}
	for i, v := range values {
		if expected[i] != v {
			t.Fatalf("expected got ss[%d] to be %s but got: %s", i, expected[i], v)
		}
	}
}

func TestSSWithSquareBrackets(t *testing.T) {
	var ss stringSliceValue
	f := setUpSSFlagSet(&ss)

	in := []string{`"[a-z]"`, `"[a-z]+"`}
	expected := []string{"[a-z]", "[a-z]+"}
	argfmt := "--ss=%s"
	arg1 := fmt.Sprintf(argfmt, in[0])
	arg2 := fmt.Sprintf(argfmt, in[1])
	err := f.Parse([]string{arg1, arg2})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	values := ss.GetSlice()

	if len(expected) != len(values) {
		t.Fatalf("expected number of values to be %d but got: %d", len(expected), len(values))
	}
	for i, v := range values {
		if expected[i] != v {
			t.Fatalf("expected got ss[%d] to be %s but got: %s", i, expected[i], v)
		}
	}
}

func TestSSAsSliceValue(t *testing.T) {
	var ss stringSliceValue
	f := setUpSSFlagSet(&ss)

	in := []string{"one", "two"}
	argfmt := "--ss=%s"
	arg1 := fmt.Sprintf(argfmt, in[0])
	arg2 := fmt.Sprintf(argfmt, in[1])
	err := f.Parse([]string{arg1, arg2})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	f.VisitAll(func(f *pflag.Flag) {
		if val, ok := f.Value.(pflag.SliceValue); ok {
			_ = val.Replace([]string{"three"})
		}
	})
	getSS := ss.GetSlice()
	if len(getSS) != 1 || getSS[0] != "three" {
		t.Fatalf("Expected ss to be overwritten with 'three', but got: %s", getSS)
	}
}
