package include

import (
	"os"
	"strings"
)

func CheckAllFilesOnly(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	filesStr := make([]string, len(files))
	for i := 0; i < len(files); i++ {
		info := CheckNameAndExt(files[i])
		filesStr[i] = info
	}

	return filesStr, nil
}

func CheckUnhiddenFilesOnly(path string) ([]string, error) {

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var unHiddenFiles []string

	for i, name := range files {
		if !strings.HasPrefix(name.Name(), ".") {
			info := CheckNameAndExt(files[i])
			unHiddenFiles = append(unHiddenFiles, info)
		}
	}
	return unHiddenFiles, nil
}
