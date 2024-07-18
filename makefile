.PHONY: run
run:
	@echo "Running the program..."
	go run eight.go 1 2 3 4 5 6 7 8

.PHONY: run-laum
run-laum:
	@echo "Running the program..."
	go run eight.go หมอ ลำ นี่ มา ตั้ง แต่ เวียง จันทน์

.PHONY: build-windows-arm
build-windows-arm:
	@echo "Building the program for Windows ARM..."
	GOOS=windows GOARCH=arm go build -o eight-arm.exe eight.go

.PHONY: build-windows-amd64
build-windows-amd64:
	@echo "Building the program for Windows AMD64..."
	GOOS=windows GOARCH=amd64 go build -o eight-amd64.exe eight.go

.PHONY: build-linux-arm
build-linux-arm:
	@echo "Building the program for Linux ARM..."
	GOOS=linux GOARCH=arm go build -o eight-arm eight.go

.PHONY: build-linux-amd64
build-linux-amd64:
	@echo "Building the program for Linux AMD64..."
	GOOS=linux GOARCH=amd64 go build -o eight-amd64 eight.go

.PHONY: build-macos-amd64
build-macos-amd64:
	@echo "Building the program for MacOS AMD64..."
	GOOS=darwin GOARCH=amd64 go build -o eight-amd64 eight.go

.PHONY: build-macos-arm64
build-macos-arm64:
	@echo "Building the program for MacOS ARM64..."
	GOOS=darwin GOARCH=arm64 go build -o eight-arm64 eight.go

.PHONY: build
build: build-windows-arm build-windows-amd64 build-linux-arm build-linux-amd64 build-macos-amd64 build-macos-arm64
	@echo "Building the program for all platforms..."