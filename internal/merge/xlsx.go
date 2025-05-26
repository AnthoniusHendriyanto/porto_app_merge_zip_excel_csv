package merge

import (
	"github.com/xuri/excelize/v2"
)

func processXLSX(path string, mergedFile *excelize.File, sheet string, currentRow *int, writeHeader *bool) error {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return err
	}
	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
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
