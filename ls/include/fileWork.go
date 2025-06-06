package include

import (
	"fmt"
	"os"
)

func GetAll(path string, flags []string) ([]string, error) {

	for i := 0; i < len(flags); i++ {
		switch flags[i] {
		case "-a":
			files, err := CheckAllFilesOnly(path)
			if err != nil {
				return nil, err
			}
			return files, nil

		case "-l":
			files, err := CheckUnhiddenFiles(path)
			if err != nil {
				return nil, err
			}

			info, err := CheckInfo(path, files)
			if err != nil {
				return nil, err
			}
			return conv(info, files), nil

		case "-la", "-al":
			files, err := CheckAllFiles(path)
			if err != nil {
				return nil, err
			}

			info, err := CheckInfo(path, files)
			if err != nil {
				return nil, err
			}
			return conv(info, files), nil
		}
	}

	return CheckUnhiddenFilesOnly(path)
}

func conv(arrStr []string, arrFiles []os.DirEntry) []string {
	strS := make([]string, len(arrStr))

	for i := 0; i < len(arrStr); i++ {
		strS[i] = fmt.Sprintf("%s %s", arrStr[i], CheckNameAndExt(arrFiles[i]))
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
