package merge

import (
	"encoding/csv"
	"os"

	"github.com/xuri/excelize/v2"
)

func processCSV(path string, mergedFile *excelize.File, sheet string, currentRow *int, writeHeader *bool) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}
	for i, row := range rows {
		if i == 0 && !*writeHeader {
			continue
		}
		for col, val := range row {
			cell, _ := excelize.CoordinatesToCellName(col+1, *currentRow)
			if err := mergedFile.SetCellValue(sheet, cell, val); err != nil {
				return err
			}
		}
		*currentRow++
	}
	*writeHeader = false
	return nil
}
