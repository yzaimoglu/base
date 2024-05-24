ifneq (,$(wildcard ./.env))
    include .env
    export
endif
TARGET = ./bin/${NAME}

.PHONY: run dev prod backend-install tailwind tailwind-build tailwind-install templ templ-dev templ-build install build dev

run:
	wgo -file=.go -file=.templ -xfile=_templ.go npx tailwindcss -i ./ui/input.css -o ./ui/static/main.css :: templ generate :: go run main.go serve
	
dev:
	@npx tailwindcss -i ./ui/input.css -o ./ui/static/main.css :: templ generate :: go run main.go serve

prod: build
	@$(TARGET) serve

backend-install:
	@go mod tidy

tailwind:
	@npx tailwindcss -i ./ui/input.css -o ./ui/static/main.css --watch

tailwind-build:
	@npx tailwindcss -i ./ui/input.css -o ./ui/static/main.css

tailwind-install:

templ: templ-dev
	
templ-dev:
	@templ generate --watch

templ-build:
	@templ fmt .
	@templ generate

backend-build:
	@go build -o $(TARGET) main.go 
	
dev: tailwind-build templ-build

install: tailwind-build backend-install

build: tailwind-build templ-build backend-build
	
clean:
	@rm -f $(TARGET)
