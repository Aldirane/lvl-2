package rune_unpack

import (
	"errors"
	"fmt"
	"testing"
)

func TestUnpackChars(t *testing.T) {
	input := "a4bc2d5e"
	want := "aaaabccddddde"
	msg, err := UnpackString(input)
	if want != msg || err != nil {
		t.Fatalf(`Result = %q, %v, want match for %#q`, msg, err, want)
	}
}

func TestUnpackCharsLong(t *testing.T) {
	input := "a10bc2d5e"
	want := "aaaaaaaaaabccddddde"
	msg, err := UnpackString(input)
	if want != msg || err != nil {
		t.Fatalf(`Result = %q, %v, want match for %#q`, msg, err, want)
	}
}

func TestSimpleString(t *testing.T) {
	input := "abcd"
	want := "abcd"
	msg, err := UnpackString(input)
	if want != msg || err != nil {
		t.Fatalf(`Result = %q, %v, want match for %#q`, msg, err, want)
	}
}

func TestEmpty(t *testing.T) {
	input := ""
	want := ""
	msg, err := UnpackString(input)
	if want != msg || err != nil {
		t.Fatalf(`Result = %q, %v, want match for %#q`, msg, err, want)
	}
}

func TestErrorDigitEscape(t *testing.T) {
	input := "45"
	want := errors.New("digit is not escaped")
	msg, err := UnpackString(input)
	if fmt.Sprint(want) != fmt.Sprint(err) || err == nil {
		t.Fatalf(`Result = %q, %#q, want match for %#q`, msg, err, want)
	}
}

func TestDigitEscape(t *testing.T) {
	input := "qwe\\4\\5"
	want := "qwe45"
	msg, err := UnpackString(input)
	if want != msg || err != nil {
		t.Fatalf(`Result = %q, %v, want match for %#q`, msg, err, want)
	}
}

func TestDigitEscapeDigit(t *testing.T) {
	input := "qwe\\45"
	want := "qwe44444"
	msg, err := UnpackString(input)
	if want != msg || err != nil {
		t.Fatalf(`Result = %q, %v, want match for %#q`, msg, err, want)
	}
}

func TestDigitEscapeEscape(t *testing.T) {
	input := "qwe\\\\5"
	want := "qwe\\\\\\\\\\"
	msg, err := UnpackString(input)
	if want != msg || err != nil {
		t.Fatalf(`Result = %q, %v, want match for %#q`, msg, err, want)
	}
}

func TestDigitEscapeEscapeDigit(t *testing.T) {
	input := "qwe\\\\4\\5e"
	want := "qwe\\\\\\\\5e"
	msg, err := UnpackString(input)
	if want != msg || err != nil {
		t.Fatalf(`Result = %q, %v, want match for %#q`, msg, err, want)
	}
}
