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
	var repeatCount = 1
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
				err = errors.New("digit is not escaped")
				return err
			}
		}
		if *targetChar == 0 {
			*targetChar, _ = utf8.DecodeRune([]byte(""))
			*repeatCount = 0
		}
		writeStr(strBuild, r, prevChar, targetChar, escapedChar, escaped, repeatCount)
		return nil
	}
	if unicode.IsDigit(r) {
		if unicode.IsDigit(*prevChar) {
			if !*escaped {
				if unicode.IsDigit(*targetChar) {
					err = errors.New("digit is not escaped")
				} else {
					*repeatCount = addInt(*repeatCount, r)
					*prevChar = r
				}
			} else {
				*repeatCount = addInt(*repeatCount, r)
				*prevChar = r
			}
		} else {
			if string(*targetChar) == "\\" && *first == false {
				*targetChar = r
				*prevChar = rune(1)
			} else {
				*first = false
				*prevChar = r
				*repeatCount, _ = strconv.Atoi(string(*prevChar))
			}
		}
	} else {
		if string(r) == "\\" {
			if string(*targetChar) != "\\" && string(*prevChar) != "\\" {
				*first = false
			} else if string(*targetChar) == "\\" && string(*prevChar) != "\\" {
				*first = false
			} else {
				*first = true
			}
			if !*first {
				writeStr(strBuild, r, prevChar, targetChar, escapedChar, escaped, repeatCount)
			}
			*escaped = true
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

func addInt(repeatCount int, r rune) int {
	intRune, _ := strconv.Atoi(string(r))
	repeatCount = repeatCount*10 + intRune
	return repeatCount
}
