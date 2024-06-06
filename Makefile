ifneq (,$(wildcard ./.env))
    include .env
    export
endif
TARGET = ./bin/${NAME}

.PHONY: run dev prod backend-install tailwind tailwind-build tailwind-install install build dev

run:
	wgo -file=.go npm run build :: go run main.go serve
	
dev:
	@npm run build :: go run main.go serve

prod: build
	@$(TARGET) serve

backend-install:
	@go install github.com/bokwoon95/wgo@latest
	@go mod tidy

tailwind:
	@npm run build-watch

tailwind-build:
	@npm run build

tailwind-install:
	@npm install

backend-build:
	@go build -o $(TARGET) main.go 
	
dev: tailwind-build

install: tailwind-install tailwind-build backend-install
	@echo "Installed dependencies."
	
build: tailwind-build backend-build
	
clean:
	@rm -f $(TARGET)
