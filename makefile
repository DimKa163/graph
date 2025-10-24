# Имя бинарного файла
BINARY_NAME=graph

# Путь для установки бинарника
BIN_DIR=bin

CMD_DIR=cmd

# Флаги сборки
BUILD_FLAGS=-ldflags "-s -w"

# Стандартные цели
.PHONY: all build run test clean lint fmt tidy help

# Основная цель
all: build


## build: Собрать бинарный файл
build:
	@echo "🔨 Building $(BINARY_NAME)..."
	@mkdir -p $(BIN_DIR)
	@go build $(BUILD_FLAGS) -o $(BIN_DIR)/$(BINARY_NAME) ./$(CMD_DIR)

## run: Запустить приложение
run:
	@echo "Running..."
	@go run ./$(CMD_DIR)

## test: Запустить тесты с покрытием
test:
	@echo "Running tests..."
	@go test -v -cover ./...

## lint: Проверить код линтером (golangci-lint)
lint:
	@echo "Linting..."
	@golangci-lint run

## fmt: Отформатировать код
fmt:
	@echo "Formatting..."
	@go fmt ./...

## tidy: Обновить go.mod/go.sum
tidy:
	@echo "Tidying..."
	@go mod tidy

## clean: Очистить бинарники
clean:
	@echo "Cleaning..."
	@rm -rf $(BIN_DIR)

## help: Показать доступные команды
help:
	@echo "Available commands:"
	@grep -E '^##' Makefile | sed -e 's/## //'
