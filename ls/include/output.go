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

func OutputFlagC(files []string) {
	width := getTerminalWidth()
	maxLen := maxFileNameLength(files)
	var cols int

	cols = (width + 2) / (maxLen + 2)
	if cols <= 0 {
		cols = 1
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer w.Flush()

	for i, file := range files {
		fmt.Fprintf(w, "%s\t", file)
		if (i+1)%cols == 0 {
			fmt.Fprintln(w)
		}
	}

	if len(files)%cols != 0 {
		fmt.Fprintln(w)
	}
}

func (d *DirCont) OutputFlagA(path string) error {
	info, err := d.CheckAllFilesOnly(path)
	if err != nil {
		return err
	}

	OutputFlagC(info)
	return nil
}
