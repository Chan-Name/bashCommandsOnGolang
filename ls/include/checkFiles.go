package include

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CheckAllFiles(path string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (d *DirCont) CheckNameAndExt(dirEntry os.DirEntry) string {

	if dirEntry.IsDir() {
		return (fmt.Sprintf("🗁 %s", dirEntry.Name()))
	}

	switch filepath.Ext(dirEntry.Name()) {
	case ".mp3", ".flac", ".wav", ".m4a", ".ogg", ".aac", ".alac", ".opus":
		return (fmt.Sprintf("♫ %s", dirEntry.Name()))

	case ".mp4", ".avi", ".mkv", ".mov", ".wmv", ".flv", ".webm":
		return (fmt.Sprintf(" %s", dirEntry.Name()))

	case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".tiff":
		return (fmt.Sprintf(" %s", dirEntry.Name()))

	case ".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".txt":
		return (fmt.Sprintf(" %s", dirEntry.Name()))

	case ".zip", ".rar", ".7z", ".tar", ".gz":
		return (fmt.Sprintf(" %s", dirEntry.Name()))

	default:
		return (fmt.Sprintf("🖹 %s", dirEntry.Name()))

	}
}

func (d *DirCont) ToUnhiddenFiles() error {
	var unHiddenFiles []os.DirEntry

	for _, name := range *d.dir {
		if !strings.HasPrefix(name.Name(), ".") {
			unHiddenFiles = append(unHiddenFiles, name)
		}
	}

	*d.dir = unHiddenFiles
	return nil
}

func (d *DirCont) CheckInfo(path string) ([]string, error) {

	fullInfo := make([]string, len(*d.dir))

	var fullPath string

	for i, v := range *d.dir {
		fullPath = fmt.Sprintf("%s/%s", path, v.Name())

		pers, err := os.Stat(fullPath)

		fullInfo[i] = fmt.Sprintf("%s, %s, %s",
			pers.Mode(), checkSize(pers.Size()),
			pers.ModTime().Format("Mon Jan _2 15:04:05 2006"))
		if err != nil {
			return nil, err
		}
	}

	return fullInfo, nil
}

func checkSize(size int64) string {
	switch {
	case size >= 1024 && size < 1048576:
		return fmt.Sprint(size/1024, "KB")

	case size >= 1048576 && size < 1073741824:
		return fmt.Sprint(size/(1024*1024), "MB")

	case size >= 1073741824 && size < 1099511627776:
		return fmt.Sprint(size/(1024*1024*1024), "GB")

	case size >= 1099511627776:
		return fmt.Sprint(size/(1024*1024*1024*1024), "TB")

	default:
		return fmt.Sprint(size, "B")

	}
}
