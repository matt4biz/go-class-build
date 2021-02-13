SOURCES := $(wildcard *.go cmd/*/*.go)

VERSION=$(shell git describe --tags --long --dirty 2>/dev/null)

## we must have tagged the repo at least once for VERSION to work
ifeq ($(VERSION),)
	VERSION = UNKNOWN
endif

sort: $(SOURCES)
	go build -ldflags "-X main.version=${VERSION}" -o $@ ./cmd/sort

.PHONY: lint
lint:
	golangci-lint run

.PHONY: committed
committed:
	@git diff --exit-code > /dev/null || (echo "** COMMIT YOUR CHANGES FIRST **"; exit 1)

docker: $(SOURCES) build/Dockerfile
	sed -e "/FIXME/s/FIXME/${VERSION}/" < build/Dockerfile|docker build -t sort-anim:latest . -f -

.PHONY: publish
publish: committed lint
	make docker
	docker tag  sort-anim:latest matthol2/sort-anim:$(VERSION)
	docker push matthol2/sort-anim:$(VERSION)
