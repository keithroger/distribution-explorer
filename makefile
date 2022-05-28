

build:
	npm run build

run:
	npm run build && go run .

lint:
	golangcli-lint run --enable-all

docker:
	docker build --tag server .

docker-run:
	docker run -p 8080:80 server

docker-dev:
	docker run -p 8080:8080 server