package main

import "fmt"

type TablePrinter struct {
	lineLength   int
	totalColumns int
	header       []string
	rows         [][]string
}

func NewTablePrinter(lineLength int, header []string, rows [][]string) *TablePrinter {
	return &TablePrinter{
		totalColumns: len(header),
		lineLength:   lineLength,
		header:       header,
		rows:         rows,
	}
}

func (tp *TablePrinter) PrintTable() {
	tp.printHeader()
	tp.printBody()
}

func (tp *TablePrinter) printHeader() {
	tp.printTableLine()
	tp.printLineContent(tp.header)
}

func (tp *TablePrinter) printBody() {
	tp.printTableLine()

	for _, lineContent := range tp.rows {
		tp.printLineContent(lineContent)
		tp.printTableLine()
	}
}

func (tp *TablePrinter) printTableLine() {
	emptyString := ""
	for i := 0; i < tp.lineLength; i++ {
		emptyString += "-"
	}
	fmt.Println(emptyString)
}

func (tp *TablePrinter) printLineContent(lineContent []string) {
	rangeToIncrease := tp.lineLength / tp.totalColumns
	partialRow := ""
	partiaLength := 0

	headerRow := ""

	for _, column := range lineContent {
		partialRow += "| "
		partiaLength += len(partialRow)

		partialRow += column
		partiaLength += len(column)

		baseSizeValue := rangeToIncrease - partiaLength
		for i := 0; i < baseSizeValue; i++ {
			partialRow += " "
			partiaLength += 1
		}

		headerRow += partialRow

		partialRow = ""
		partiaLength = 0
	}
	headerRow += "|"
	fmt.Println(headerRow)
}

func main() {
	tablePrinter := NewTablePrinter(
		100,
		[]string{"task", "status", "started", "initDate", "endDate"},
		[][]string{
			{"something to do", "wip", "true", "2021-01-1", "2021-01-01"},
			{"something to do", "wip", "true", "2021-01-1", "2021-01-01"},
			{"something to do", "wip", "true", "2021-01-1", "2021-01-01"},
			{"something to do", "wip", "true", "2021-01-1", "2021-01-01"},
			{"something to do", "wip", "true", "2021-01-1", "2021-01-01"},
			{"something to do", "wip", "true", "2021-01-1", "2021-01-01"},
		},
	)

	tablePrinter.PrintTable()
}
