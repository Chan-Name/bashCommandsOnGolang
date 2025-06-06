package main

import (
	"gol/ls/include"
	"log"
)

func main() {
	path, flags := include.GetLaunchOptions()
	a, err := include.GetAll(path, flags)
	if err != nil {
		log.Fatal("ERROR", err)
	}

	cols := include.CalculateColumns(a, 2)
	include.PrettyPrintColumns(a, cols)

}
