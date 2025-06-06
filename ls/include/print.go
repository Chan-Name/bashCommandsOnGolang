package include

import (
	"fmt"
	"os"
	"text/tabwriter"

	"golang.org/x/term"
)

func getTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 80
	}
	return width
}

func maxFileNameLength(files []string) int {
	maxLen := 0
	for _, file := range files {
		if len(file) > maxLen {
			maxLen = len(file)
		}
	}
	return maxLen
}

func CalculateColumns(files []string, padding int) int {
	width := getTerminalWidth()
	maxLen := maxFileNameLength(files)

	if maxLen == 0 {
		return 1
	}

	cols := (width + padding) / (maxLen + padding)
	if cols <= 0 {
		return 1
	}
	return cols
}

func PrettyPrintColumns(files []string, columns int) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer w.Flush()

	for i, file := range files {
		fmt.Fprintf(w, "%s\t", file)
		if (i+1)%columns == 0 {
			fmt.Fprintln(w)
		}
	}

	if len(files)%columns != 0 {
		fmt.Fprintln(w)
	}
}
