p=$(shell pwd)

test.unit.service:
	PROJ_DIR=$p go test -count=1 -mod=vendor -v ./pkg/service

test.all:
	test.unit.service

server:
	PROJ_DIR=$p GOFLAGS=-mod=vendor go run ./main.go server

migration:
	PROJ_DIR=$p GOFLAGS=-mod=vendor go run ./main.go migration

SWAGGER_FILE := documents/swagger.json
API_HEADER_FILE := $(p)/internal/pkg/delivery/restful/router.go
API_PATH := $(p)/internal/pkg
swagger.gen:
	# go get -u github.com/mikunalpha/goas
	goas --module-path . --main-file-path $(API_HEADER_FILE) --handler-path $(API_PATH) --output $(SWAGGER_FILE)
	make swagger.copy

swagger.server:
	docker run -d --rm --name sg -p 8088:8080 -e SWAGGER_JSON=/documents/swagger.json -v $(p)/documents:/documents swaggerapi/swagger-ui

swagger.copy: 
	cp ./documents/swagger.json /Users/zh/Documents/github/swagger/swagger-ui/dev-helpers/swagger.json