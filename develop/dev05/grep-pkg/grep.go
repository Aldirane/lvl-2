package grep_pkg

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func Command(args []string) int {
	var grepApp GrepApp
	err := grepApp.setFlags(args)
	if err != nil {
		return 2
	}

	if err = grepApp.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	return 0
}

type GrepApp struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
	Pattern    string
	reader     *os.File
	Input      []string
}

func (grepApp *GrepApp) setFlags(args []string) error {
	fl := flag.NewFlagSet("grep_pkg", flag.ContinueOnError)
	fl.IntVar(&grepApp.After, "A", 0, "Печатать +N строк после совпадения")
	fl.IntVar(&grepApp.Before, "B", 0, "Печатать +N строк до совпадения")
	fl.IntVar(&grepApp.Context, "C", 0, "Печатать ±N строк вокруг совпадения")
	fl.BoolVar(&grepApp.Count, "c", false, "Количество строк")
	fl.BoolVar(&grepApp.IgnoreCase, "i", false, "Игнорировать регистр")
	fl.BoolVar(&grepApp.Invert, "v", false, "Вместо совпадения, исключать")
	fl.BoolVar(&grepApp.Fixed, "F", false, "Точное совпадение со строкой, не паттерн")
	fl.BoolVar(&grepApp.LineNum, "n", false, "Печатать номер строки")

	if err := fl.Parse(args); err != nil {
		//fl.Usage()
		return err
	}

	// Проверка наличия аргумента поиска
	if len(fl.Arg(0)) == 0 {
		err := errors.New("Не указано слово для поиска")
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}

	grepApp.Pattern = fl.Arg(0)
	if grepApp.Fixed {
		grepApp.Pattern = fmt.Sprintf("^%s$", grepApp.Pattern)
	}

	if grepApp.IgnoreCase {
		grepApp.Pattern = strings.ToLower(grepApp.Pattern)
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		grepApp.reader = os.Stdin
		return nil
	}

	file, err := os.Open(fl.Arg(1))
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't open file %s: %v\n", fl.Arg(0), err)
		return err
	}
	grepApp.reader = file

	return nil
}

func (grepApp *GrepApp) run() error {
	defer grepApp.reader.Close()
	scanner := bufio.NewScanner(grepApp.reader)
	for scanner.Scan() {
		line := scanner.Text()
		if grepApp.IgnoreCase {
			line = strings.ToLower(line)
		}
		grepApp.Input = append(grepApp.Input, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if grepApp.Count {
		CountLines := grepApp.CountLines()
		fmt.Printf("Total lines match = %d\n", CountLines)
	} else {
		toPrintLines := grepApp.GetLines()
		for _, line := range toPrintLines {
			fmt.Print(line)
		}
	}

	return nil
}

// Функция для поиска и печати строк с указанными параметрами
func (grepApp *GrepApp) GetLines() []string {
	BeforeLines := grepApp.Input[:]
	AfterLines := grepApp.Input[:]
	ContextLines := grepApp.Input[:]
	toPrint := make([]string, 0)
	for lineCount, line := range grepApp.Input {

		// Проверка наличия совпадения
		match := false
		if grepApp.Fixed {
			match = (line == grepApp.Pattern)
		} else {
			match = strings.Contains(line, grepApp.Pattern)
		}

		// Печать строки с учетом параметров
		if match && !grepApp.Invert {
			if grepApp.LineNum {
				toPrint = append(toPrint, fmt.Sprintf("match at %d: ", lineCount))
			}
			toPrint = append(toPrint, fmt.Sprintln(line))

			// Печать строк до совпадения
			if grepApp.Before > 0 {
				toPrint = append(toPrint, fmt.Sprintln("Before match"))
				start := lineCount - grepApp.Before
				if start < 0 {
					start = 0
				}
				for i := start; i < lineCount; i++ {
					bl := BeforeLines[i]
					match = strings.Contains(bl, grepApp.Pattern)
					if !match {
						if grepApp.LineNum {
							toPrint = append(toPrint, fmt.Sprintf("index %d: ", i))
						}
						toPrint = append(toPrint, fmt.Sprintln(bl))
					}
				}
			}

			// Печать строк после совпадения
			if grepApp.After > 0 {
				toPrint = append(toPrint, fmt.Sprintln("After match"))
				start := lineCount + 1
				finish := lineCount + grepApp.After + 1
				if finish > len(grepApp.Input) {
					finish = len(grepApp.Input)
				}
				for i := start; i < finish; i++ {
					al := AfterLines[i]
					match = strings.Contains(al, grepApp.Pattern)
					if !match {
						if grepApp.LineNum {
							toPrint = append(toPrint, fmt.Sprintf("index %d: ", i))
						}
						toPrint = append(toPrint, fmt.Sprintln(al))
					}
				}
			}

			// Печать строк вокруг совпадения
			if grepApp.Context > 0 {
				toPrint = append(toPrint, fmt.Sprintln("Context lines"))
				start := lineCount - grepApp.Context
				finish := lineCount + grepApp.Context + 1
				if start < 0 {
					start = 0
				}
				if finish > len(grepApp.Input) {
					finish = len(grepApp.Input)
				}
				for i := start; i < finish; i++ {
					cl := ContextLines[i]
					match = strings.Contains(cl, grepApp.Pattern)
					if !match {
						if grepApp.LineNum {
							toPrint = append(toPrint, fmt.Sprintf("index %d: ", i))
						}
						toPrint = append(toPrint, fmt.Sprintln(cl))
					}
				}
			}
		} else if !match && grepApp.Invert {
			if grepApp.LineNum {
				toPrint = append(toPrint, fmt.Sprintf("invert match at %d: ", lineCount))
			}
			toPrint = append(toPrint, fmt.Sprintln(line))

			// Печать строк до совпадения
			if grepApp.Before > 0 {
				toPrint = append(toPrint, fmt.Sprintln("Before invert match"))
				start := lineCount - grepApp.Before
				if start < 0 {
					start = 0
				}
				for i := start; i < lineCount; i++ {
					bl := BeforeLines[i]
					match = strings.Contains(bl, grepApp.Pattern)
					if !match {
						if grepApp.LineNum {
							toPrint = append(toPrint, fmt.Sprintf("index %d: ", i))
						}
						toPrint = append(toPrint, fmt.Sprintln(bl))
					}
				}
			}

			// Печать строк после совпадения
			if grepApp.After > 0 {
				toPrint = append(toPrint, fmt.Sprintln("After match"))
				for i := lineCount + 1; i < len(grepApp.Input); i++ {
					al := AfterLines[i]
					match = strings.Contains(al, grepApp.Pattern)
					if !match {
						if grepApp.LineNum {
							toPrint = append(toPrint, fmt.Sprintf("index %d: ", i))
						}
						toPrint = append(toPrint, fmt.Sprintln(al))
					}
				}
			}

			// Печать строк вокруг совпадения
			if grepApp.Context > 0 {
				toPrint = append(toPrint, fmt.Sprintln("Context lines"))
				start := lineCount - grepApp.Context
				finish := lineCount + grepApp.Context + 1
				if start < 0 {
					start = 0
				}
				if finish > len(grepApp.Input) {
					finish = len(grepApp.Input)
				}
				for i := start; i < finish; i++ {
					cl := ContextLines[i]
					match = strings.Contains(cl, grepApp.Pattern)
					if !match {
						if grepApp.LineNum {
							toPrint = append(toPrint, fmt.Sprintf("index %d: ", i))
						}
						toPrint = append(toPrint, fmt.Sprintln(cl))
					}
				}
			}
		}
	}

	return toPrint
}

// Функция для подсчета количества строк с указанным параметром
func (grepApp *GrepApp) CountLines() int {
	count := 0
	for _, line := range grepApp.Input {
		// Проверка наличия совпадения
		match := false
		match = strings.Contains(line, grepApp.Pattern)
		if match {
			count++
		}
	}
	return count
}
