/*
Создать Go-функцию, осуществляющую примитивную распаковку
строки, содержащую повторяющиеся символы/руны, например:
● "a4bc2d5e" => "aaaabccddddde"
● "abcd" => "abcd"
● "45" => "" (некорректная строка)
● "" => ""
Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
● qwe\4\5 => qwe45 (*)
● qwe\45 => qwe44444 (*)
● qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка, функция
должна возвращать ошибку. Написать unit-тесты.
*/

package rune_unpack

import (
	"errors"
	"strconv"
	"strings"
)

func UnpackString(input string) (string, error) {
	var result strings.Builder
	var repeatCount int
	escaped := false

	for _, char := range input {
		if escaped {
			result.WriteRune(char)
			escaped = false
			continue
		}

		if char == '\\' {
			escaped = true
			continue
		}

		if char >= '0' && char <= '9' {
			repeatCount, _ = strconv.Atoi(string(char))
		} else {
			if repeatCount == 0 {
				repeatCount = 1
			}
			result.WriteString(strings.Repeat(string(char), repeatCount))
			repeatCount = 0
		}
	}

	if escaped {
		return "", errors.New("invalid input string")
	}

	return result.String(), nil
}
