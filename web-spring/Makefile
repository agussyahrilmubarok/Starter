.PHONY: help dev build test clean run jar check format lint deps update kill

APP_NAME=web-spring

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
	./mvnw spring-boot:run

build:
	./mvnw clean package

test:
	./mvnw test

clean:
	./mvnw clean

jar:
	./mvnw package

run:
	java -jar target/*.jar

check:
	./mvnw verify

format:
	./mvnw spotless:apply

lint:
	./mvnw spotless:check

deps:
	./mvnw dependency:tree

update:
	./mvnw versions:display-dependency-updates

kill:
	lsof -ti:8080 | xargs kill -9 || true