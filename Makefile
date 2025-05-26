# Makefile to build zip-merge-tool for Linux and Windows

APP_NAME = zip-merge-tool
SRC_DIR = ./cmd/gui

.PHONY: all clean linux windows

all: linux windows

linux:
	GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME)-linux $(SRC_DIR)

windows:
	GOOS=windows GOARCH=amd64 go build -o bin/$(APP_NAME)-windows.exe $(SRC_DIR)

clean:
	rm -rf bin
