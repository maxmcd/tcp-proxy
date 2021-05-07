
.PHONY: build push

build: README.md
	docker build -t maxmcd/tcp-proxy .

push: build
	docker push maxmcd/tcp-proxy

README.md: ./main.go ./tools/*.md
	./tools/build_readme.sh

