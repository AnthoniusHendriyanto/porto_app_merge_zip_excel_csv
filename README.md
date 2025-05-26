# ZIP to Merged Excel Tool (GUI)

A simple cross-platform desktop GUI tool written in Go using the Fyne framework. It allows you to merge multiple `.csv` and `.xlsx` files from a `.zip` archive into a single Excel file.

---

## ✅ Features

- Select a `.zip` archive containing `.csv` and `.xlsx` files
- Merge files into a single `.xlsx` with headers included only once
- Auto-increment output filename if it already exists (e.g., `merged_1.xlsx`)
- Choose destination folder for output
- Cancel merging mid-process
- Automatically opens the output folder after merging
- Clean, portable GUI (no browser required)

---

## 🖥️ Installation & Build

> Requires Go 1.18+ and Git. Docker is **not required** if you build from Windows.

### 1. Clone the Project
```bash
git clone https://github.com/yourname/merge-zip-gui.git
cd merge-zip-gui
```

### 2. Initialize Go Modules
```bash
go mod tidy
```

### 3. Build for Windows (from Windows machine)
Use the Makefile:
```bash
make windows
```
Or manually:
```bash
go build -o zip-merge-tool.exe ./cmd/gui
```

### 4. Build for Linux (from Linux machine)
Use the Makefile:
```bash
make linux
```
Or manually:
```bash
go build -o zip-merge-tool ./cmd/gui
```bash
go build -o zip-merge-tool ./cmd/gui
```

> If you are on Linux and want to build Windows version, you will need `fyne-cross` with Docker.

---

## 🗂 Folder Structure

```
merge-zip-gui/
├── cmd/
│   └── gui/
│       └── main.go
├── internal/
│   ├── merge/
│   │   ├── processor.go
│   │   ├── csv.go
│   │   ├── xlsx.go
│   │   ├── unzip.go
│   │   └── open.go
│   └── ui/
│       └── layout.go
├── Makefile
├── go.mod
├── LICENSE
└── README.md
```

---

## 📦 Usage

1. Launch the app by running the compiled binary:
   ```bash
   ./zip-merge-tool.exe  # On Windows
   ./zip-merge-tool      # On Linux
   ```
2. Select a `.zip` file that contains `.csv` or `.xlsx`
3. Select a destination folder
4. Click **Merge and Save**

✅ Done! The merged Excel file will appear in your selected folder.

---

## 🛠 Dependencies

- [Fyne](https://github.com/fyne-io/fyne) — GUI Framework
- [Excelize](https://github.com/qax-os/excelize) — Excel Reader/Writer

---

## 📄 License

MIT License © 2025 Anthonius Hendriyanto