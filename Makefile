ifneq (,$(wildcard ./.env))
    include .env
    export
endif
TARGET = ./bin/${NAME}

.PHONY: run dev prod backend-install tailwind tailwind-build tailwind-install templ templ-dev templ-build install build dev

run:
	wgo -file=.go -file=.templ -xfile=_templ.go npm run build :: templ generate :: go run main.go serve
	
dev:
	@npm run build :: templ generate :: go run main.go serve

prod: build
	@$(TARGET) serve

backend-install:
	@go install github.com/bokwoon95/wgo@latest
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get github.com/a-h/templ
	@go mod tidy

tailwind:
	@npm run build-watch

tailwind-build:
	@npm run build

tailwind-install:
	@npm install

templ: templ-dev
	
templ-dev:
	@templ generate --watch

templ-build:
	@templ fmt .
	@templ generate

backend-build:
	@go build -o $(TARGET) main.go 
	
dev: tailwind-build templ-build

install: tailwind-install tailwind-build backend-install
	@echo "Installed dependencies."
	
build: tailwind-build templ-build backend-build
	
clean:
	@rm -f $(TARGET)
