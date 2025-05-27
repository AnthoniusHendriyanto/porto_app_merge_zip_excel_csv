# Makefile to build zip-merge-tool for Linux and Windows

APP_NAME = zip-merge-tool
SRC_DIR = ./cmd/gui

.PHONY: all clean linux windows

all: linux windows

linux:
	GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME)-linux $(SRC_DIR)

windows:
	go build -o bin/zip-merge-tool.exe ./cmd/gui

clean:
	rm -rf bin
