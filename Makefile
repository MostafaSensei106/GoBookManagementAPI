# Makefile for building GoWebBase
# Author: Mostafa Sensei106
# License: MIT
#
# Note for Windows Users:
# This Makefile uses POSIX shell commands (like 'rm', 'cp', 'mkdir -p').
# For best results on Windows, please run 'make' commands from a
# POSIX-compliant shell like Git Bash or within WSL/MSYS2.

# declare variables
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
APP_NAME := GoBookManagementAPI
OUTPUT_DIR := bin/$(GOOS)/$(GOARCH)
OUTPUT := $(OUTPUT_DIR)/$(APP_NAME)
GOWEBBASE_VERSION := 1.0.0

.PHONY: all build clean release help check deps fmt vet install docker-build docker-run

all: build

deps:
	 @echo "📦 Checking dependencies..."
	 @if [ -f go.sum ]; then \
		echo "📦 Verifying dependencies..."; \
		go mod verify; \
		echo "✅ Dependencies installed and up-to-date"; \
	else \
		echo "📦 Downloading dependencies..."; \
		go mod download; \
		echo "📦 Verifying dependencies..."; \
		go mod verify; \
		echo "✅ Dependencies installed"; \
	fi

fmt:
	 @echo "🎨 Formatting code..."
	 @go fmt ./...

vet:
	 @echo "🔎 Vetting code..."
	 @go vet ./...

check: deps fmt vet

build: check
	 @echo "📦 Building $(APP_NAME) for $(GOOS)/$(GOARCH)..."
	 @mkdir -p $(OUTPUT_DIR)
	 @GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(OUTPUT) .
	 @echo "✅ Build complete: $(OUTPUT)"

install: build
	@echo "✅ $(APP_NAME) built successfully. Find the executable in the '$(OUTPUT_DIR)' directory."
	@echo "This project is a web service and is not installed to a system-wide directory."

release: check
	 @{ \
		echo "🔍 Detecting host platform..."; \
		HOST_OS=$$(go env GOOS); \
		HOST_ARCH=$$(go env GOARCH); \
		echo "🖥️  Host: $$HOST_OS/$$HOST_ARCH"; \
		echo "🌐 Building for all major platforms and architectures..."; \
		platforms="linux/386 linux/amd64 linux/arm linux/arm64 windows/386 windows/amd64 windows/arm windows/arm64"; \
		for platform in $$platforms; do \
			GOOS=$${platform%/*}; \
			GOARCH=$${platform#*/}; \
			OUT_DIR=bin/$$GOOS/$$GOARCH; \
			OUT_FILE=$$OUT_DIR/$(APP_NAME); \
			if [ "$$GOOS" = "windows" ]; then \
				OUT_FILE=$$OUT_FILE.exe; \
			fi; \
			ARCHIVE_NAME=$(APP_NAME)-v$(GOWEBBASE_VERSION)-$$GOOS-$$GOARCH; \
			mkdir -p $$OUT_DIR; \
			echo "🛠️  Building for $$GOOS/$$GOARCH..."; \
			if [ "$$GOOS" = "windows" ] && [ "$$HOST_OS" != "windows" ] && ! command -v x86_64-w64-mingw32-gcc >/dev/null 2>&1; then \
				echo "⚠️ Skipped: $$GOOS/$$GOARCH (Windows cross-compiler 'x86_64-w64-mingw32-gcc' not found)"; \
				continue; \
			fi; \
			GOOS=$$GOOS GOARCH=$$GOARCH go build -o $$OUT_FILE . || { echo "❌ Build Failed for $$GOOS/$$GOARCH"; continue; }; \
			echo "✅ Build done: $$OUT_FILE"; \
			mkdir -p release; \
			if [ "$$GOOS" = "windows" ]; then \
				(cd bin && zip -r "../release/$$ARCHIVE_NAME.zip" "$$GOOS/$$GOARCH" >/dev/null) && \
				echo "✅ Compressed (zip): release/$$ARCHIVE_NAME.zip"; \
			else \
				(cd bin && tar -czf "../release/$$ARCHIVE_NAME.tar.gz" "$$GOOS/$$GOARCH" >/dev/null) && \
				echo "✅ Compressed (tar.gz): release/$$ARCHIVE_NAME.tar.gz"; \
			fi; \
		done; \
		echo "🎉 Release archives created in the 'release' directory."; \
	}

docker-build:
	@echo "🐳 Building Docker image..."
	@docker build -t $(APP_NAME):latest .
	@echo "✅ Docker image '$(APP_NAME):latest' built successfully."

docker-run: docker-build
	@echo "🚀 Running Docker container..."
	@docker run -p 8080:8080 $(APP_NAME):latest

clean:
	 @echo "🧹 Cleaning build artifacts..."
	 @rm -rf bin release
	 @go clean -cache -modcache -testcache
	 @echo "✅ Clean complete."

help:
	 @echo ""
	 @echo "📖 GoBookManagementAPI Makefile Commands"
	 @echo "============================="
	 @echo "make all           👉 Alias for 'make build'."
	 @echo "make deps          👉 Check and download Go module dependencies."
	 @echo "make fmt           👉 Format all Go source files."
	 @echo "make vet           👉 Run 'go vet' to check for suspicious constructs."
	 @echo "make check         👉 Run all checks (deps, fmt, vet)."
	 @echo "make build         👉 Build the 'gowebbase' executable for the current OS/architecture."
	 @echo "make install       👉 An alias for 'make build'. Does not install system-wide."
	 @echo "make release       👉 Build and package for all target platforms (Linux, Windows)."
	 @echo "make docker-build  👉 Build the Docker image for the application."
	 @echo "make docker-run    👉 Build and run the application inside a Docker container."
	 @echo "make clean         👉 Delete all build artifacts, release archives, and Go caches."
	 @echo "make help          👉 Show this help message."
	 @echo ""
