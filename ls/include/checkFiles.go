package include

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CheckNameAndExt(dirEntry os.DirEntry) string {

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

func CheckAllFiles(path string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func CheckUnhiddenFiles(path string) ([]os.DirEntry, error) {
	var unHiddenFiles []os.DirEntry

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, name := range files {
		if !strings.HasPrefix(name.Name(), ".") {
			unHiddenFiles = append(unHiddenFiles, name)
		}
	}
	return unHiddenFiles, nil
}

func CheckInfo(path string, files []os.DirEntry) ([]string, error) {

	fullInfo := make([]string, len(files))

	var fullPath string

	for i, v := range files {
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
