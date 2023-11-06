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

package main

import unpack "unpack-pkg/rune-unpack"

func main() {
	input := "\\4\\5"
	result, err := unpack.UnpackString(input)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println(result)
	}
	input = "a10bc2d5e2"
	result, err = unpack.UnpackString(input)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println(result)
	}
	input = "qwe\\5\\4e"
	result, err = unpack.UnpackString(input)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println(result)
	}
	input = "qwe\\\\4e"
	result, err = unpack.UnpackString(input)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println(result)
	}
}
