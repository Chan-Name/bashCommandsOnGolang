package include

import (
	"os"
	"strings"
)

func (d *DirCont) CheckAllFilesOnly(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	filesStr := make([]string, len(files))
	for i := 0; i < len(files); i++ {
		info := d.CheckNameAndExt((*d.dir)[i])
		filesStr[i] = info
	}

	return filesStr, nil
}

func (d *DirCont) CheckUnhiddenFilesOnly(path string) ([]string, error) {

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var unHiddenFiles []string

	for i, name := range files {
		if !strings.HasPrefix(name.Name(), ".") {
			info := d.CheckNameAndExt((*d.dir)[i])
			unHiddenFiles = append(unHiddenFiles, info)
		}
	}
	return unHiddenFiles, nil
}
