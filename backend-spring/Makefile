.PHONY: help dev build test clean run jar check format lint deps update kill

APP_NAME=backend-spring

help:
	@echo "Available commands:"
	@echo "  make dev       - Run Spring Boot in development mode"
	@echo "  make build     - Build project"
	@echo "  make test      - Run tests"
	@echo "  make clean     - Clean build files"
	@echo "  make run       - Run built jar"
	@echo "  make jar       - Generate jar file"
	@echo "  make check     - Run all verification"
	@echo "  make deps      - Show dependencies"
	@echo "  make update    - Refresh dependencies"
	@echo "  make kill      - Kill app running on 8080"

dev:
	./gradlew bootRun

build:
	./gradlew build

test:
	./gradlew test

clean:
	./gradlew clean

jar:
	./gradlew bootJar

run:
	java -jar build/libs/$(APP_NAME).jar

check:
	./gradlew check

format:
	./gradlew spotlessApply

lint:
	./gradlew spotlessCheck

deps:
	./gradlew dependencies

update:
	./gradlew --refresh-dependencies

kill:
	lsof -ti:8080 | xargs kill -9 || true