ifndef GOROOT
	export GOROOT=$(realpath $(CURDIR)/../go)
	export PATH := $(GOROOT)/bin:$(PATH)
endif
ifndef GOPATH
	export GOPATH=$(CURDIR)/frontend
	export PATH := $(GOPATH)/bin:$(PATH)
endif


ALL_GO_SOURCES=$(shell /bin/sh -c "find *.go | grep -v _test.go")

docker:
	sudo docker-compose up --build

docker-local-postgres:
	sudo docker build --tag url-shortener .
	sudo docker-compose -f docker-compose-local-postgres.yml up

context-default:
	sudo docker context use default

context-esc:
	sudo docker context use myecscontext

check_compose:
	docker-compose config

ls_image:
	sudo docker image ls

ls_volume:
	sudo docker volume list

ls_network:
	sudo docker network list

ls_context:
	sudo docker context ls

docker_ps:
	sudo docker ps

docker_ps_all:
	sudo docker ps -all

run: fmt
	go run $(ALL_GO_SOURCES)

js: fmt
	@echo "\033[92mGenerating JavaScript ...\033[0m"
	@gopherjs build frontend/*.go -m -o frontend/app.js

test_curl_create1:
	curl -XPOST -d'{"url":"https://github.com/siongui/go-kit-url-shortener-micro-service"}' localhost:8080/create
test_curl_create2:
	curl -XPOST -d'{"url":"https://github.com/siongui/goef"}' localhost:8080/create

test_curl_metrics:
	curl -XGET localhost:8080/metrics

test: fmt
	cd datasource; go test -v -race

fmt:
	go fmt *.go
	go fmt frontend/*.go

modinit:
	go mod init github.com/siongui/go-kit-url-shortener-micro-service

modtidy:
	go mod tidy -compat=1.17

install_gopherjs:
	go get -u github.com/gopherjs/gopherjs
