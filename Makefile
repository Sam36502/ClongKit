PROJECT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
DIST_DIR := ${PROJECT_DIR}dist/
DIST_LIN := lin/
INSTALLDIR_LIN := /usr/bin/
DIST_WIN := win/
EXE_NAME := clongkit


.PHONY: help build build-lin build-win

help: ## Display this help menu
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: build-lin build-win ## Build executable binaries for all supported systems

build-lin: ## Build for GNU+Linux
	@echo ' > Building for Linux'
	@GOOS=linux go build -o ${DIST_DIR}${DIST_LIN}${EXE_NAME} ./*.go

build-win: ## Build for Windows
	@echo ' > Building for Windows'
	@GOOS=windows go build -o ${DIST_DIR}${DIST_WIN}${EXE_NAME}.exe ./*.go

install-lin: build-lin ## Install on a GNU+Linux system
	@echo ' > Installing to ${INSTALLDIR_LIN}...'
	@sudo ln -s ${DIST_DIR}${DIST_LIN}${EXE_NAME} ${INSTALLDIR_LIN}${EXE_NAME}

uninstall-lin: ## Uninstall on a GNU+Linux system
	@echo ' > Removing ${INSTALLDIR_LIN}${EXE_NAME}...'
	@sudo rm -f ${INSTALLDIR_LIN}${EXE_NAME}
