NAME = base
TARGET = ./${NAME}

.PHONY: run clean build fdev bdev fbuild bbuild prod install finstall binstall

run: build
	@go run main.go serve

prod: build
	@./$(TARGET) serve

fdev:
	@npm run dev --prefix ./ui

bdev:
	@go run base.go serve

fbuild:
	@npm run build --prefix ./ui

bbuild:
	@go build -o $(TARGET) main.go

finstall:
	@npm install --prefix ./ui

binstall:
	@go mod tidy

install: finstall binstall

build: fbuild bbuild
	
clean:
	@rm -f $(TARGET)
