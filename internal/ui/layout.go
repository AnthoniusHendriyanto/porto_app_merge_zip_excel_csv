package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"porto_app_merge_zip_excel_csv/internal/merge"
)

var cancelFlag atomic.Bool

func LaunchApp() {
	a := app.New()
	w := a.NewWindow("ZIP to Merged Excel")
	status := widget.NewLabel("Please upload a .zip file containing .csv/.xlsx")

	zipPathEntry := widget.NewEntry()
	zipPathEntry.Disable()
	outputDirEntry := widget.NewEntry()
	outputDirEntry.Disable()
	outputFileEntry := widget.NewEntry()
	outputFileEntry.SetText("merged.xlsx")

	cancelButton := widget.NewButton("Cancel", func() {
		cancelFlag.Store(true)
		status.SetText("Cancelled by user")
	})
	cancelButton.Disable()

	progress := dialog.NewCustomWithoutButtons("Processing", container.NewVBox(
		widget.NewLabel("Merging, please wait..."),
		cancelButton,
	), w)

	selectZip := widget.NewButton("1. Select ZIP File", func() {
		dlg := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil || reader == nil {
				return
			}
			path := reader.URI().Path()
			if !strings.HasSuffix(strings.ToLower(path), ".zip") {
				dialog.ShowError(fmt.Errorf("only .zip files are allowed"), w)
				return
			}
			zipPathEntry.SetText(path)
		}, w)
		dlg.SetFilter(storage.NewExtensionFileFilter([]string{".zip"}))
		dlg.Show()
	})

	selectOutputDir := widget.NewButton("2. Choose Output Folder", func() {
		dialog.NewFolderOpen(func(folder fyne.ListableURI, err error) {
			if err != nil || folder == nil {
				return
			}
			outputDirEntry.SetText(folder.Path())
		}, w).Show()
	})

	runButton := widget.NewButton("3. Merge and Save", func() {
		zipPath := zipPathEntry.Text
		outputDir := outputDirEntry.Text
		filename := strings.TrimSpace(outputFileEntry.Text)

		if zipPath == "" || outputDir == "" {
			dialog.ShowError(fmt.Errorf("ZIP file and output folder must be selected"), w)
			return
		}

		if filename == "" {
			filename = "merged.xlsx"
		} else if !strings.HasSuffix(strings.ToLower(filename), ".xlsx") {
			filename += ".xlsx"
		}

		finalFilename := filename
		counter := 1
		for {
			outputPath := filepath.Join(outputDir, finalFilename)
			if _, err := os.Stat(outputPath); os.IsNotExist(err) {
				break
			}
			base := strings.TrimSuffix(filename, ".xlsx")
			finalFilename = fmt.Sprintf("%s_%d.xlsx", base, counter)
			counter++
		}

		outputPath := filepath.Join(outputDir, finalFilename)
		status.SetText("Processing: " + filepath.Base(zipPath))
		cancelFlag.Store(false)
		cancelButton.Enable()
		progress.Show()

		go func() {
			err := merge.ProcessZip(zipPath, outputPath, &cancelFlag)
			time.Sleep(300 * time.Millisecond)
			progress.Hide()
			cancelButton.Disable()
			if cancelFlag.Load() {
				status.SetText("Merge cancelled.")
				return
			}
			if err != nil {
				dialog.ShowError(err, w)
				status.SetText("Error: " + err.Error())
			} else {
				dialog.ShowInformation("Success", "Merged data saved to "+outputPath, w)
				status.SetText("Done: " + outputPath)
				zipPathEntry.SetText("")
				outputDirEntry.SetText("")
				outputFileEntry.SetText("merged.xlsx")
				merge.OpenFolder(outputDir)
			}
		}()
	})

	form := container.NewVBox(
		widget.NewLabel("Merge .csv and .xlsx from a .zip file into one Excel sheet"),
		selectZip,
		zipPathEntry,
		selectOutputDir,
		outputDirEntry,
		widget.NewLabel("Output Filename (e.g. merged.xlsx):"),
		outputFileEntry,
		runButton,
		status,
	)

	w.SetContent(form)
	w.Resize(fyne.NewSize(520, 380))
	w.ShowAndRun()
}
