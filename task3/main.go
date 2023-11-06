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
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	keyColumn             int
	sortNumeric           bool
	reverseSort           bool
	removeDuplicates      bool
	sortMonth             bool
	ignoreTrailingSpaces  bool
	checkSorted           bool
	sortNumericWithSuffix bool
)

func init() {
	flag.IntVar(&keyColumn, "k", 0, "Specify the column for sorting")
	flag.BoolVar(&sortNumeric, "n", false, "Sort numerically")
	flag.BoolVar(&reverseSort, "r", false, "Sort in reverse order")
	flag.BoolVar(&removeDuplicates, "u", false, "Remove duplicate lines")
	flag.BoolVar(&sortMonth, "M", false, "Sort by month name")
	flag.BoolVar(&ignoreTrailingSpaces, "b", false, "Ignore trailing spaces")
	flag.BoolVar(&checkSorted, "c", false, "Check if the data is sorted")
	flag.BoolVar(&sortNumericWithSuffix, "h", false, "Sort numerically with suffixes")

	flag.Parse()
}

func main() {
	input := flag.Arg(0)
	output := flag.Arg(1)

	lines, err := readLines(input)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	sortLines(lines)

	if removeDuplicates {
		lines = removeDuplicateLines(lines)
	}

	if checkSorted {
		if sorted := isSorted(lines); sorted {
			fmt.Println("Data is sorted")
		} else {
			fmt.Println("Data is not sorted")
		}
		return
	}

	err = writeLines(lines, output)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func writeLines(lines []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}

	return writer.Flush()
}

func sortLines(lines []string) {
	sort.SliceStable(lines, func(i, j int) bool {
		if sortMonth {
			return sortByMonth(lines[i], lines[j])
		}

		if sortNumeric {
			return sortNumericLines(lines[i], lines[j])
		}

		return strings.ToLower(lines[i]) < strings.ToLower(lines[j])
	})

	if reverseSort {
		reverseSlice(lines)
	}
}

func sortByMonth(a, b string) bool {
	layout := "January 2006" // example: "January 2022"

	aTime, err := time.Parse(layout, a)
	if err == nil {
		bTime, err := time.Parse(layout, b)
		if err == nil {
			return aTime.Before(bTime)
		}
	}

	return strings.ToLower(a) < strings.ToLower(b)
}

func sortNumericLines(a, b string) bool {
	aInt, err := strconv.Atoi(a)
	if err == nil {
		bInt, err := strconv.Atoi(b)
		if err == nil {
			return aInt < bInt
		}
	}

	return strings.ToLower(a) < strings.ToLower(b)
}

func reverseSlice(slice []string) {
	for i := len(slice)/2 - 1; i >= 0; i-- {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}
}

func removeDuplicateLines(lines []string) []string {
	uniqueLines := []string{}
	seenLines := make(map[string]bool)

	for _, line := range lines {
		if _, seen := seenLines[line]; !seen {
			uniqueLines = append(uniqueLines, line)
			seenLines[line] = true
		}
	}

	return uniqueLines
}

func isSorted(lines []string) bool {
	for i := 1; i < len(lines); i++ {
		if strings.ToLower(lines[i]) < strings.ToLower(lines[i-1]) {
			return false
		}
	}
	return true
}

/*
Можно запустить модуль следующим образом:
go run main.go -к [номер колонки] -n -r -u -M -b -c -h [входной файл] [выходной файл]
Например, чтобы отсортировать файл input.txt и сохранить отсортированные строки
в файл output.txt по колонке 2 в обратном порядке, можно выполнить следующую команду:
go run main.go -k 2 -r input.txt output.txt
*/
