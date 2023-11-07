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

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	rune_unpack "unpack-pkg/rune-unpack"
)

func main() {
	var input = ""
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.Trim(input, "\n")
	result, err := rune_unpack.UnpackString(input)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	} else {
		fmt.Println(result)
	}
}
