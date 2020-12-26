package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rows := []uint{128, 128, 128, 128}
	var err error
	for true {
		rows, err = GenerateRow(rows)
		if err != nil {
			panic(err)
		}
		for _, r := range rows {
			fmt.Printf("%s\n", RenderLine(r))
		}
		fmt.Print("Press 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
