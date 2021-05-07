
.PHONY: build push

build:
	docker build -t maxmcd/tcp-proxy .

push: build
	docker push maxmcd/tcp-proxy
