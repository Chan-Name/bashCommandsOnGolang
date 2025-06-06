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

	for i := 0; i < len(a); i++ {
		include.PrintInfo(a[i])
	}

}
