package sort_pkg

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

type SortModule struct {
	keyColumn             int
	sortNumeric           bool
	reverseSort           bool
	removeDuplicates      bool
	sortMonth             bool
	ignoreTrailingSpaces  bool
	checkSorted           bool
	sortNumericWithSuffix bool
	input                 *os.File
	output                *os.File
}

// Command line runs the go-sort command line app and returns its exit status.
func Command(args []string) int {
	var sortApp SortModule
	err := sortApp.setFlags(args)
	if err != nil {
		return 2
	}
	if err = sortApp.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	return 0
}

func (sortApp *SortModule) setFlags(args []string) error {
	sortFlag := flag.NewFlagSet("sort module", flag.ContinueOnError)
	sortFlag.IntVar(&sortApp.keyColumn, "k", 1, "Specify the column for sorting")
	sortFlag.BoolVar(&sortApp.sortNumeric, "n", false, "Sort numerically")
	sortFlag.BoolVar(&sortApp.reverseSort, "r", false, "Sort in reverse order")
	sortFlag.BoolVar(&sortApp.removeDuplicates, "u", false, "Remove duplicate lines")
	sortFlag.BoolVar(&sortApp.sortMonth, "M", false, "Sort by month name")
	sortFlag.BoolVar(&sortApp.ignoreTrailingSpaces, "b", false, "Ignore trailing spaces")
	sortFlag.BoolVar(&sortApp.checkSorted, "c", false, "Check if the data is sorted")
	sortFlag.BoolVar(&sortApp.sortNumericWithSuffix, "h", false, "Sort numerically with suffixes")

	if err := sortFlag.Parse(args); err != nil {
		flag.Usage()
		return err
	}
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		sortApp.input = os.Stdin
		sortApp.output = os.Stdout
		return nil
	}
	input, err := os.OpenFile(sortFlag.Arg(0), os.O_RDONLY, 0444)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't open file %s: %v\n", sortFlag.Arg(0), err)
		return err
	}
	output, err := os.OpenFile(sortFlag.Arg(1), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("can't open/create file %s: %v\n", sortFlag.Arg(0), err)
		output = os.Stdout
	}
	sortApp.input = input
	sortApp.output = output
	return nil
}

func (sortApp *SortModule) run() error {
	defer sortApp.input.Close()
	lines, err := sortApp.readLines()
	if err != nil {
		return err
	}
	lines = sortApp.sortLines(lines)

	if sortApp.checkSorted {
		if sorted := sortApp.isSorted(lines); sorted {
			fmt.Println("Data is sorted")
		} else {
			fmt.Println("Data is not sorted")
		}
		return nil
	}

	if sortApp.keyColumn == 1 && !sortApp.sortNumeric {
		lines = sortApp.sortLines(lines)
		sortApp.writeLines(lines)
		return nil
	}

	lines = sortApp.sortColumns(lines)
	err = sortApp.writeLines(lines)
	if err != nil {
		return err
	}
	return nil
}

func (sortApp *SortModule) readLines() ([]string, error) {
	defer sortApp.input.Close()
	scanner := bufio.NewScanner(sortApp.input)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func (sortApp *SortModule) writeLines(lines []string) error {
	defer sortApp.output.Close()
	writer := bufio.NewWriter(sortApp.output)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}

	return writer.Flush()
}

func (sortApp *SortModule) sortLines(lines []string) []string {
	sort.SliceStable(lines, func(i, j int) bool {
		if sortApp.sortMonth {
			return sortApp.sortByMonth(lines[i], lines[j])
		}

		if sortApp.sortNumeric {
			return sortApp.sortNumericLines(lines[i], lines[j])
		}

		return strings.ToLower(lines[i]) < strings.ToLower(lines[j])
	})
	if sortApp.removeDuplicates {
		lines = sortApp.removeDuplicateLines(lines)
	}
	if sortApp.reverseSort {
		sortApp.reverseSlice(lines)
	}
	return lines
}

func (sortApp *SortModule) sortByMonth(a, b string) bool {
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

func (sortApp *SortModule) sortNumericLines(a, b string) bool {
	aInt, err := strconv.Atoi(a)
	if err == nil {
		bInt, err := strconv.Atoi(b)
		if err == nil {
			return aInt < bInt
		}
	}
	return strings.ToLower(a) < strings.ToLower(b)
}

func (sortApp *SortModule) reverseSlice(slice []string) {
	for i := len(slice)/2 - 1; i >= 0; i-- {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}
}

func (sortApp *SortModule) removeDuplicateLines(lines []string) []string {
	uniqueLines := make([]string, 0)
	seenLines := make(map[string]bool)
	for _, line := range lines {
		if _, seen := seenLines[line]; !seen {
			uniqueLines = append(uniqueLines, line)
			seenLines[line] = true
		}
	}
	return uniqueLines
}

func (sortApp *SortModule) isSorted(lines []string) bool {
	for i := 1; i < len(lines); i++ {
		if strings.ToLower(lines[i]) < strings.ToLower(lines[i-1]) {
			return false
		}
	}
	return true
}

func (sortApp *SortModule) sortColumns(lines []string) []string {
	row := sortRowByColumn{
		data:        make([][]string, 0, len(lines)),
		keyColumn:   sortApp.keyColumn - 1,
		sortNumeric: sortApp.sortNumeric,
	}

	for _, v := range lines {
		row.data = append(row.data, strings.Fields(v))
	}

	if sortApp.reverseSort {
		sort.Sort(sort.Reverse(row))
	} else {
		sort.Sort(row)
	}

	for i, v := range row.data {
		lines[i] = strings.Join(v, " ")
	}

	if sortApp.removeDuplicates {
		lines = sortApp.removeDuplicateLines(lines)
	}

	return lines
}
