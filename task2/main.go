package main

import unpack "unpack-pkg/rune-unpack"

func main() {
	// Пример использования
	input := "a4bc2d5e"
	result, err := unpack.UnpackString(input)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println(result)
	}
}
