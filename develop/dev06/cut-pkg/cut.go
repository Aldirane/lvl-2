package cut_pkg

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type CutApp struct {
	fields    string
	delimiter string
	separated bool
	reader    *os.File
}

func Command(args []string) int {
	var cutApp CutApp
	err := cutApp.setFlags(args)
	if err != nil {
		return 2
	}

	if err = cutApp.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	return 0
}

func (cutApp *CutApp) setFlags(args []string) error {
	fl := flag.NewFlagSet("cut_pkg", flag.ContinueOnError)
	fl.StringVar(&cutApp.fields, "f", "", "выбрать поля (колонки)")
	fl.StringVar(&cutApp.delimiter, "d", "\t", "использовать другой разделитель")
	fl.BoolVar(&cutApp.separated, "s", false, "только строки с разделителем")
	if err := fl.Parse(args); err != nil {
		fl.Usage()
		return err
	}
	cutApp.reader = os.Stdin
	return nil
}

func (cutApp *CutApp) run() error {
	// Чтение строк из STDIN
	scanner := bufio.NewScanner(cutApp.reader)
	for scanner.Scan() {
		line := scanner.Text()

		// Если включен флаг -s, пропускаем строки без разделителя
		if cutApp.separated && !strings.Contains(line, cutApp.delimiter) {
			continue
		}

		// Разбиваем строку на колонки
		columns := strings.Split(line, cutApp.delimiter)

		// Если указан флаг -f, выбираем только указанные колонки
		if cutApp.fields != "" {
			fieldIndexes := cutApp.parseFields(cutApp.fields)
			selectedColumns := make([]string, len(fieldIndexes))
			for i, index := range fieldIndexes {
				if index >= 0 && index < len(columns) {
					selectedColumns[i] = columns[index]
				}
			}
			fmt.Println(strings.Join(selectedColumns, cutApp.delimiter))
		} else {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения из STDIN:", err)
		return err
	}
	return nil
}

// Парсинг списка выбранных полей
func (cutApp *CutApp) parseFields(fields string) []int {
	fieldIndexes := []int{}
	fieldRanges := strings.Split(fields, ",")
	for _, fieldRange := range fieldRanges {
		if strings.Contains(fieldRange, "-") {
			rangeParts := strings.Split(fieldRange, "-")
			start := cutApp.parseInt(rangeParts[0])
			end := cutApp.parseInt(rangeParts[1])
			if start < 0 || end < 0 || start > end {
				continue
			}
			for i := start; i <= end; i++ {
				fieldIndexes = append(fieldIndexes, i-1)
			}
		} else {
			index := cutApp.parseInt(fieldRange)
			if index > 0 {
				fieldIndexes = append(fieldIndexes, index-1)
			}
		}
	}
	return fieldIndexes
}

// Преобразование строки в число
func (cutApp *CutApp) parseInt(str string) int {
	var result int
	_, err := fmt.Sscanf(str, "%d", &result)
	if err != nil {
		return -1
	}
	return result
}
