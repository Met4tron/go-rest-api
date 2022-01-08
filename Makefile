BINARY=build

dev:
	go run ./cmd/go-rest-api/main.go

build:
	go build -o ${BINARY} ./cmd/go-rest-api/main.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

run:
	docker-compose up --build -d

build-docker:
	docker build -t go-rest-api .

.PHONY: clean install build docker run