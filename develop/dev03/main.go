/*
Утилита sort
Отсортировать строки в файле по аналогии с консольной
утилитой sort (man sort — смотрим описание и основные
параметры): на входе подается файл из несортированными
строками, на выходе — файл с отсортированными.
Реализовать поддержку утилитой следующих ключей:
-k — указание колонки для сортировки (слова в строке могут
выступать в качестве колонок, по умолчанию разделитель —
пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки
Дополнительно
Реализовать поддержку утилитой следующих ключей:
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учетом суффиксов
*/

package main

import (
	sort_pkg "mysort/sort-pkg"
	"os"
)

func main() {
	os.Exit(sort_pkg.Command(os.Args[1:]))
}

/*
Можно запустить модуль следующим образом:
go run main.go -к [номер колонки] -n -r -u -M -b -c -h [входной файл] [выходной файл]
Например, чтобы отсортировать файл input.txt и сохранить отсортированные строки
в файл output.txt по колонке 2 в обратном порядке, можно выполнить следующую команду:
go run main.go -k 2 -r input.txt output.txt
*/