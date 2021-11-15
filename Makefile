ifndef GOROOT
	export GOROOT=$(realpath $(CURDIR)/../go)
	export PATH := $(GOROOT)/bin:$(PATH)
endif


ALL_GO_SOURCES=$(shell /bin/sh -c "find *.go | grep -v _test.go")

docker:
	sudo docker-compose up --build

check_compose:
	docker-compose config

ls_image:
	sudo docker image ls

ls_volume:
	sudo docker volume list

ls_network:
	sudo docker network list

docker_ps:
	sudo docker ps

docker_ps_all:
	sudo docker ps -all

run: fmt
	go run $(ALL_GO_SOURCES)

test_curl_create1:
	curl -XPOST -d'{"url":"https://github.com/siongui/go-kit-url-shortener-micro-service"}' localhost:8080/create
test_curl_create2:
	curl -XPOST -d'{"url":"https://github.com/siongui/goef"}' localhost:8080/create

test_curl_metrics:
	curl -XGET localhost:8080/metrics

test: fmt
	go test -v -race

fmt:
	go fmt *.go

modinit:
	go mod init github.com/siongui/go-kit-url-shortener-micro-service

modtidy:
	go mod tidy -compat=1.17
