package rune_unpack

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func UnpackString(input string) (string, error) {
	var err error
	var strBuild strings.Builder
	var result = ""
	var byteStr = []byte(input)
	var repeatCount int = 1
	var targetChar, prevChar, escapedChar rune
	escaped := false
	first := false
	for utf8.RuneCount(byteStr) >= 0 {
		if utf8.RuneCount(byteStr) == 0 {
			err = escapeChar(
				&strBuild,
				rune(0),
				&prevChar,
				&targetChar,
				&escapedChar,
				&escaped,
				&first,
				&repeatCount)
			break
		}
		r, size := utf8.DecodeRune(byteStr)
		byteStr = byteStr[size:]
		if string(r) == "\\" {
			escaped = true
		}
		if prevChar == 0 {
			targetChar = r
			prevChar = r
			continue
		}
		err = escapeChar(
			&strBuild,
			r,
			&prevChar,
			&targetChar,
			&escapedChar,
			&escaped,
			&first,
			&repeatCount)
		if err != nil {
			break
		}
	}
	result = strBuild.String()
	return result, err
}

func escapeChar(
	strBuild *strings.Builder,
	r rune,
	prevChar, targetChar, escapedChar *rune,
	escaped, first *bool,
	repeatCount *int,
) error {
	var err error
	if r == 0 {
		if unicode.IsDigit(*targetChar) {
			if !*escaped {
				err = errors.New("Digit is not escaped")
				return err
			}
		}
		writeStr(strBuild, r, prevChar, targetChar, escapedChar, escaped, repeatCount)
	}
	if unicode.IsDigit(r) {
		if unicode.IsDigit(*prevChar) {
			if !*escaped {
				if unicode.IsDigit(*targetChar) {
					err = errors.New("Digit is not escaped")
				} else {
					*repeatCount = buidInt(*repeatCount, r)
					*prevChar = r
				}
			} else {
				*repeatCount = buidInt(*repeatCount, r)
				*prevChar = r
			}
		} else { // "\\4\\5"
			if string(*targetChar) == "\\" && *first == false {
				*targetChar = r
				*prevChar = rune(1)
			} else {
				*prevChar = r
				*repeatCount, _ = strconv.Atoi(string(*prevChar))
			}
		}
	} else {
		if string(r) == "\\" {
			if string(*targetChar) != "\\" && string(*prevChar) != "\\" {
				*first = false
			} else {
				*first = true
			}
			writeStr(strBuild, r, prevChar, targetChar, escapedChar, escaped, repeatCount)
			*escaped = true
			*targetChar = r
			*prevChar = r
		} else {
			writeStr(strBuild, r, prevChar, targetChar, escapedChar, escaped, repeatCount)
		}
	}
	return err
}

func writeStr(
	strBuild *strings.Builder,
	r rune,
	prevChar, targetChar, escapedChar *rune,
	escaped *bool,
	repeatCount *int,
) {
	*escapedChar = *targetChar
	*prevChar = r
	*targetChar = *prevChar
	strBuild.WriteString(strings.Repeat(string(*escapedChar), *repeatCount))
	*repeatCount = 1
	*escaped = false
}

func buidInt(repeatCount int, r rune) int {
	intr, _ := strconv.Atoi(string(r))
	repeatCount = repeatCount*10 + intr
	return repeatCount
}
