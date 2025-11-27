# 项目名称
PROJECT_NAME := bill-server

# 主入口文件
MAIN_FILE := ./cmd/server/main.go

# 输出目录
OUTPUT_DIR := ./bin

# 输出二进制文件名
OUTPUT_BIN := $(OUTPUT_DIR)/$(PROJECT_NAME)

# Go 编译参数
GO_BUILD_FLAGS := -ldflags "-s -w"

# 环境变量
ENV := development

.PHONY: all build run clean test deps

# 默认目标
all: build

# 编译项目
build:
	@mkdir -p $(OUTPUT_DIR)
	@echo "Building $(PROJECT_NAME)..."
	@go build $(GO_BUILD_FLAGS) -o $(OUTPUT_BIN) $(MAIN_FILE)
	@echo "Build completed: $(OUTPUT_BIN)"

# 运行项目
run:
	@echo "Running $(PROJECT_NAME)..."
	@go run $(MAIN_FILE)

# 清理构建产物
clean:
	@echo "Cleaning up..."
	@rm -rf $(OUTPUT_DIR)
	@echo "Clean completed"

# 运行测试
test:
	@echo "Running tests..."
	@go test -v ./...

# 安装依赖
deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@echo "Dependencies installed"

# 生成 API 文档
docs:
	@echo "Generating API docs..."
	@swag init -g $(MAIN_FILE)
	@echo "API docs generated"
