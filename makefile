help:        ## Help command
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

build:       ## Build with npm
	@npm run build

run:         ## Build with npm and run go server
	npm run build && go run .

lint:        ## Linting
	@golangcli-lint run --enable-all

docker:      ## Build docker container
	docker build --tag server .

docker-run:  ## Run docker container at local port 80
	docker run -p 80:8080 server

docker-dev:  ## Run docker container at local port 8080
	docker run -p 8080:8080 server