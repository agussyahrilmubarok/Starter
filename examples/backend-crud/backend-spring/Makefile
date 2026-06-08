.PHONY: help dev build test clean run jar check format lint deps update kill

APP_NAME := backend-spring
PORT := 8080
GRADLE := ./gradlew

help:
	@echo "Available commands:"
	@echo "  make dev       - Run Spring Boot in development mode"
	@echo "  make build     - Clean and build project"
	@echo "  make test      - Run tests"
	@echo "  make clean     - Clean build files"
	@echo "  make jar       - Generate executable jar"
	@echo "  make run       - Run latest generated jar"
	@echo "  make check     - Run all verification"
	@echo "  make format    - Format source code"
	@echo "  make lint      - Check code formatting"
	@echo "  make deps      - Show dependencies"
	@echo "  make update    - Refresh dependencies"
	@echo "  make kill      - Kill application running on port $(PORT)"

dev:
	$(GRADLE) bootRun

build: clean
	$(GRADLE) build

test:
	$(GRADLE) test

clean:
	$(GRADLE) clean

jar:
	$(GRADLE) bootJar

run:
	java -jar $$(ls -t build/libs/*.jar | head -n 1)

check:
	$(GRADLE) check

format:
	$(GRADLE) spotlessApply

lint:
	$(GRADLE) spotlessCheck

deps:
	$(GRADLE) dependencies

update:
	$(GRADLE) --refresh-dependencies

kill:
	@PID=$$(lsof -ti:$(PORT)); \
	if [ -n "$$PID" ]; then \
		echo "Stopping process $$PID on port $(PORT)"; \
		kill -9 $$PID; \
	else \
		echo "No process found on port $(PORT)"; \
	fi