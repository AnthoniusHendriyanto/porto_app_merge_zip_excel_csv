package merge

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync/atomic"

	"github.com/xuri/excelize/v2"
)

func ProcessZip(zipPath, outputPath string, cancelFlag *atomic.Bool) error {
	tempDir := "temp_gui"
	os.RemoveAll(tempDir)
	os.Mkdir(tempDir, 0755)

	if err := unzipFiles(zipPath, tempDir); err != nil {
		return fmt.Errorf("failed to unzip: %w", err)
	}

	mergedFile := excelize.NewFile()
	mergedFile.NewSheet("MergedData")
	mergedFile.DeleteSheet("Sheet1")

	currentRow := 1
	writeHeader := true

	err := filepath.Walk(tempDir, func(path string, info os.FileInfo, err error) error {
		if cancelFlag.Load() {
			return fmt.Errorf("cancelled")
		}
		if err != nil || info.IsDir() {
			return err
		}
		typeStr, err := detectFileType(path)
		if err != nil {
			return nil
		}
		switch typeStr {
		case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", "application/zip":
			return processXLSX(path, mergedFile, "MergedData", &currentRow, &writeHeader)
		case "text/plain", "text/csv", "application/octet-stream":
			return processCSV(path, mergedFile, "MergedData", &currentRow, &writeHeader)
		default:
			return nil
		}
	})
	if err != nil {
		return fmt.Errorf("error processing files: %w", err)
	}

	if err := mergedFile.SaveAs(outputPath); err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}
	os.RemoveAll(tempDir)
	return nil
}

func unzipFiles(zipPath, dest string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		if f.FileInfo().IsDir() {
			continue
		}
		outPath := filepath.Join(dest, filepath.Base(f.Name))
		o, err := os.Create(outPath)
		if err != nil {
			return err
		}
		rc, err := f.Open()
		if err != nil {
			o.Close()
			return err
		}
		_, err = io.Copy(o, rc)
		o.Close()
		rc.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func detectFileType(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	buf := make([]byte, 512)
	n, err := f.Read(buf)
	if err != nil && err != io.EOF {
		return "", err
	}
	return http.DetectContentType(buf[:n]), nil
}
