package include

import (
	"fmt"
	"os"
)

type DirCont struct {
	dir *[]os.DirEntry
}

func New(dir []os.DirEntry) *DirCont {
	return &DirCont{dir: &dir}
}

func (d *DirCont) GetAll(path string, flags []string) ([]string, error) {

	for i := 0; i < len(flags); i++ {
		switch flags[i] {
		case "-a":
			files, err := d.CheckAllFilesOnly(path)
			if err != nil {
				return nil, err
			}
			return files, nil

		case "-l":
			d.ToUnhiddenFiles()
			info, err := d.CheckInfo(path)
			if err != nil {
				return nil, err
			}
			return d.conv(info), nil

		case "-la", "-al":
			info, err := d.CheckInfo(path)
			if err != nil {
				return nil, err
			}
			return d.conv(info), nil
		}
	}

	return d.CheckUnhiddenFilesOnly(path)
}

func (d *DirCont) conv(arrStr []string) []string {
	strS := make([]string, len(arrStr))

	for i := 0; i < len(arrStr); i++ {
		strS[i] = fmt.Sprintf("%s %s", arrStr[i], d.CheckNameAndExt((*d.dir)[i]))
	}
	return strS
}

func GetLaunchOptions() (string, []string) {

	path, _ := os.Getwd()

	if len(os.Args) <= 1 {
		return path, nil
	}

	flags := make([]string, len(os.Args))

	for _, arg := range os.Args[1:] {
		if arg[0] == '-' && arg[1] != '-' {
			flags = append(flags, arg)
		} else {
			path = arg
		}
	}

	return path, flags
}
