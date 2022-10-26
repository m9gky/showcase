include dev.env
export

.PHONY: lint
lint:
	cd src && golangci-lint run

.PHONY: fix
fix:
	cd src && golangci-lint run --fix


import-sl: build
	./bin/showcase import-sl

start: build
	./bin/showcase start

build:
	rm -rf ./bin/showcase
	cd src && go build -o ../bin/showcase .