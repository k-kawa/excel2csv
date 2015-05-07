package main

import (
	"os"
	"fmt"
	"strings"
	"github.com/tealeg/xlsx"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("[SYNOPSIS]")
		fmt.Println("    excel2csv filename.xlxm")
		os.Exit(1)
	}

	var excelFileName = os.Args[1]
	xlFile, err := xlsx.OpenFile(excelFileName)

	if err != nil {
		fmt.Println("err :", err)
		os.Exit(1)
	}

	for _, sheet := range xlFile.Sheets {
		fmt.Printf("%s\n", sheet.Name)

		for i, row := range sheet.Rows {
			var cells []string

			for _, cell := range row.Cells {
				cells = append(cells, fmt.Sprintf("%q", cell.Formula()))
			}
			fmt.Println( fmt.Sprintf("%d", i) + ","+ strings.Join(cells, ","))
		}
	}
}
