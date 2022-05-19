

build:
	npm run build

run: build
	go run main.go

lint:
	golangcli-lint run --enable-all
