

build:
	npm run build

run:
	npm run build && go run .

lint:
	golangcli-lint run --enable-all
