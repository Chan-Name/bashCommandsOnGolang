package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	a, _ := getAllFiles(getPath())
	checkExt(a)
}

func printFiles(icon rune, name string) {
	fmt.Printf("%c %s\n", icon, name)
}

func checkExt(arr []os.DirEntry) (rune, string) {

	for i := 0; i < len(arr); i++ {

		if arr[i].IsDir() {
			printFiles('🗁', arr[i].Name())
			continue
		}

		switch filepath.Ext(arr[i].Name()) {
		case ".mp3", ".flac", ".wav", ".m4a", ".ogg", ".aac", ".alac", ".opus":
			printFiles('♫', arr[i].Name())

		case ".mp4", ".avi", ".mkv", ".mov", ".wmv", ".flv", ".webm":
			printFiles('', arr[i].Name())

		case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".tiff":
			printFiles('', arr[i].Name())

		case ".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".txt":
			printFiles('', arr[i].Name())

		case ".zip", ".rar", ".7z", ".tar", ".gz":
			printFiles('', arr[i].Name())

		default:
			printFiles('🖹', arr[i].Name())

		}
	}
	os.Exit(1)
	return 'a', ""
}

func getAllFiles(path string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func getPath() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	fmt.Println("Not enough command line arguments")
	os.Exit(1)
	return ""
}
