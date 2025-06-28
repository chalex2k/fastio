# Makefile для Go-проекта

.PHONY: install-tools lint format check-format install-linter help

# Установка всех необходимых инструментов
install-tools:
	go install mvdan.cc/gofumpt@latest
	brew install golangci-lint

# Запуск линтера
lint:
	golangci-lint run

# Форматирование всего проекта
format:
	gofumpt -w .

# Проверка форматирования (без изменений)
check-format:
	gofumpt -l . | grep -v ^$ && echo "❌ Требуется форматирование" && exit 1 || echo "✅ Все файлы отформатированы"

# Установка только golangci-lint (через brew)
install-linter:
	brew install golangci-lint

# Справка по командам
help:
	@echo "Доступные команды:"
	@echo "  make install-tools   - Установить все инструменты (gofumpt + golangci-lint)"
	@echo "  make install-linter  - Установить только golangci-lint"
	@echo "  make lint            - Запустить линтинг проекта"
	@echo "  make format          - Отформатировать весь проект"
	@echo "  make check-format    - Проверить форматирование"
	@echo "  make help            - Показать эту справку"