package main

import (
	"bytes"
	"testing"
)

func compare(t *testing.T, exp []int, got []int) {
	if len(exp) != len(got) {
		t.Fatalf("Expected:%v\nGot:%v", exp, got)
	}

	for i, v := range got {
		if v != exp[i] {

			t.Fatalf("Expected:%d\nGot:%d", exp[i], v)
		}
	}
}

func TestParseFields(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		expected := []int{0}
		got, err := parseFields("1")
		if err != nil {
			t.Error(err)
		}

		compare(t, expected, got)

	})

	t.Run("1,2", func(t *testing.T) {
		expected := []int{0, 1}
		got, err := parseFields("1,2")
		if err != nil {
			t.Error(err)
		}

		compare(t, expected, got)
	})

	t.Run("1-3", func(t *testing.T) {
		expected := []int{0, 1, 2}
		got, err := parseFields("1-3")
		if err != nil {
			t.Error(err)
		}

		compare(t, expected, got)
	})

	t.Run("1-2,4", func(t *testing.T) {
		expected := []int{0, 1, 3}
		got, err := parseFields("1-2,4")
		if err != nil {
			t.Error(err)
		}

		compare(t, expected, got)
	})

	t.Run("4,1", func(t *testing.T) {
		expected := []int{0, 3}
		got, err := parseFields("4,1")
		if err != nil {
			t.Error(err)
		}

		compare(t, expected, got)
	})

}

func TestCutUtil(t *testing.T) {
	strRows := []string{
		"zfds 310 1 xvcxv gfdgfdg",
		"cfdf 1 5 bx fds",
		"dsad 11 231 zas xvdsf",
		"aada 65 43 hgfh mfdg",
		"aada 65 432 khj khghfg",
		"aada 65 454 tret hjgfh",
		"aada 65 2 nbvnvb ergfdg",
		"aada 65 0 jhgjg jhgjhg",
		"fdsfdsf;gdfgdfds;fsdfdsf",
	}

	t.Run("test1", func(t *testing.T) {
		expected := "zfds 1\n" + "cfdf 5\n" + "dsad 231\n" + "aada 43\n" + "aada 432\n" + "aada 454\n" + "aada 2\n" + "aada 0\n" + "fdsfdsf;gdfgdfds;fsdfdsf\n"

		fields, _ := parseFields("1,3")

		cutUtil := NewCutUtil(strRows, fields, " ", false)
		var gotStr bytes.Buffer
		cutUtil.cut(&gotStr)
		if gotStr.String() != expected {
			t.Fatalf("Not equal error\nExpected:\n%s\nGot:\n%s", expected, gotStr.String())
		}
	})

	t.Run("test2", func(t *testing.T) {
		expected := "310\n" + "1\n" + "11\n" + "65\n" + "65\n" + "65\n" + "65\n" + "65\n" + "fdsfdsf;gdfgdfds;fsdfdsf\n"

		fields, _ := parseFields("2")

		cutUtil := NewCutUtil(strRows, fields, " ", false)
		var gotStr bytes.Buffer
		cutUtil.cut(&gotStr)
		if gotStr.String() != expected {
			t.Fatalf("Not equal error\nExpected:\n%s\nGot:\n%s", expected, gotStr.String())
		}
	})

	t.Run("test3", func(t *testing.T) {
		expected := "310 1 xvcxv\n" + "1 5 bx\n" + "11 231 zas\n" + "65 43 hgfh\n" + "65 432 khj\n" + "65 454 tret\n" + "65 2 nbvnvb\n" + "65 0 jhgjg\n" + "fdsfdsf;gdfgdfds;fsdfdsf\n"

		fields, _ := parseFields("2-4")

		cutUtil := NewCutUtil(strRows, fields, " ", false)
		var gotStr bytes.Buffer
		cutUtil.cut(&gotStr)
		if gotStr.String() != expected {
			t.Fatalf("Not equal error\nExpected:\n%s\nGot:\n%s", expected, gotStr.String())
		}
	})

	t.Run("test4 with separated=true", func(t *testing.T) {
		expected := "zfds 1\n" + "cfdf 5\n" + "dsad 231\n" + "aada 43\n" + "aada 432\n" + "aada 454\n" + "aada 2\n" + "aada 0\n"

		fields, _ := parseFields("1,3")

		cutUtil := NewCutUtil(strRows, fields, " ", true)
		var gotStr bytes.Buffer
		cutUtil.cut(&gotStr)
		if gotStr.String() != expected {
			t.Fatalf("Not equal error\nExpected:\n%s\nGot:\n%s", expected, gotStr.String())
		}
	})

}
