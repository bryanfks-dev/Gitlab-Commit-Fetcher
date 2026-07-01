PROJECT_NAME=gitlab-commit-fetcher
OUTPUT_DIR=./dist
GO_MAIN_FILEPATH=./cmd/main.go
OUTPUT_FILEPATH=$(OUTPUT_DIR)/$(PROJECT_NAME)

install:
	go mod tidy
	go mod download

format:
	gofmt -w . && go tool golines -m 80 -w .

dev-windows:
	GOOS=windows \
		go build \
			-o $(OUTPUT_FILEPATH) $(GO_MAIN_FILEPATH) \
			&& $(OUTPUT_FILEPATH)

dev-linux:
	GOOS=linux \
		go build \
			-o $(OUTPUT_FILEPATH) $(GO_MAIN_FILEPATH) \
			&& $(OUTPUT_FILEPATH)

dev-macos:
	GOOS=darwin \
		go build \
			-o $(OUTPUT_FILEPATH) $(GO_MAIN_FILEPATH) \
			&& $(OUTPUT_FILEPATH)

compile-linux-amd64: install
	GOOS=linux GOARCH=amd64 \
		go build \
			-o $(OUTPUT_FILEPATH)-linux-amd64 \
			$(GO_MAIN_FILEPATH)

compile-linux-arm64: install
	GOOS=linux GOARCH=arm64 \
		go build \
			-o $(OUTPUT_FILEPATH)-linux-arm64 \
			$(GO_MAIN_FILEPATH)

compile-linux: compile-linux-amd64 compile-linux-arm64

compile-macos-amd64: install
	GOOS=darwin GOARCH=amd64 \
		go build \
			-o $(OUTPUT_FILEPATH)-macos-amd64 \
			$(GO_MAIN_FILEPATH)

compile-macos-arm64: install
	GOOS=darwin GOARCH=arm64 \
		go build \
			-o $(OUTPUT_FILEPATH)-macos-arm64 \
			$(GO_MAIN_FILEPATH)

compile-macos: compile-macos-amd64 compile-macos-arm64

compile-windows-amd64: install
	GOOS=windows GOARCH=amd64 \
		go build \
			-o $(OUTPUT_FILEPATH)-macos-amd64.exe \
			$(GO_MAIN_FILEPATH)

compile-windows-arm64: install
	GOOS=windows GOARCH=arm64 \
		go build \
			-o $(OUTPUT_FILEPATH)-windows-arm64.exe \
			$(GO_MAIN_FILEPATH)

compile-windows: compile-windows-amd64 compile-windows-arm64

compile: compile-linux compile-macos compile-windows
