# Makefile for building GoBookManagementAPI
# Author: Mostafa Sensei106
# License: MIT

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
APP_NAME := GoBookManagementAPI
OUTPUT_DIR := bin/$(GOOS)/$(GOARCH)
OUTPUT := $(OUTPUT_DIR)/$(APP_NAME)

.PHONY: all build clean release help check deps fmt vet install docker-build docker-run

all: build

deps:
	@echo "📦 Checking dependencies..."
	@if [ -f go.sum ]; then \
		echo "📦 Verifying dependencies..."; \
		go mod verify; \
	else \
		echo "📦 Downloading dependencies..."; \
		go mod download; \
		echo "📦 Verifying dependencies..."; \
		go mod verify; \
	fi
	@echo "✅ Dependencies OK"

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
	@echo "✅ $(APP_NAME) built successfully in '$(OUTPUT_DIR)'"

release: check
	@echo "🌐 Building release binaries..."
	@platforms="linux/amd64 linux/arm linux/arm64 windows/amd64"; \
	for platform in $$platforms; do \
		GOOS=$${platform%/*}; \
		GOARCH=$${platform#*/}; \
		OUT_DIR=bin/$$GOOS/$$GOARCH; \
		OUT_FILE=$$OUT_DIR/$(APP_NAME); \
		[ "$$GOOS" = "windows" ] && OUT_FILE="$$OUT_FILE.exe"; \
		mkdir -p $$OUT_DIR; \
		echo "🛠️ Building for $$GOOS/$$GOARCH..."; \
		GOOS=$$GOOS GOARCH=$$GOARCH go build -o $$OUT_FILE . || { echo "❌ Build failed for $$GOOS/$$GOARCH"; continue; }; \
		echo "📦 Packaging..."; \
		ARCHIVE_NAME=$(APP_NAME)-$$GOOS-$$GOARCH; \
		mkdir -p release; \
		if [ "$$GOOS" = "windows" ]; then \
			(cd bin && zip -r "../release/$$ARCHIVE_NAME.zip" "$$GOOS/$$GOARCH" >/dev/null); \
		else \
			(cd bin && tar -czf "../release/$$ARCHIVE_NAME.tar.gz" "$$GOOS/$$GOARCH" >/dev/null); \
		fi; \
		echo "✅ Done $$GOOS/$$GOARCH"; \
	done
	@echo "🎉 Release builds are in /release"

docker-build:
	@echo "🐳 Building Docker image..."
	@docker build -t $(APP_NAME):latest .
	@echo "✅ Docker image built successfully."

docker-run:
	@echo "🚀 Running Docker container..."
	@docker run -p 8080:8080 $(APP_NAME):latest

clean:
	@echo "🧹 Cleaning..."
	@rm -rf bin release
	@go clean -cache -modcache -testcache
	@echo "✅ Clean complete."

help:
	@echo "📖 GoBookManagementAPI Makefile Commands"
	@echo "make deps          👉 Install & verify dependencies"
	@echo "make fmt           👉 Format sources"
	@echo "make vet           👉 Static analysis"
	@echo "make build         👉 Build backend only"
	@echo "make release       👉 Build release binaries for Linux/Windows"
	@echo "make docker-build  👉 Build docker image"
	@echo "make docker-run    👉 Run docker container"
	@echo "make clean         👉 Clean workspace"
