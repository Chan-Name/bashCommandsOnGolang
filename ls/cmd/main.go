package main

import (
	"gol/ls/include"
	"log"
)

func main() {
	path, flags := include.GetLaunchOptions()

	d, _ := include.CheckAllFiles(path)

	v := include.New(d)

	a, err := v.GetAll(path, flags)
	if err != nil {
		log.Fatal("ERROR", err)
	}

	cols := include.CalculateColumns(a, 2)
	include.PrettyPrintColumns(a, cols)

}
