BINARY_NAME=${HOME}/go/bin/gorunning
ifneq ($(GOPATH),)
     BINARY_NAME=${GOPATH}/bin/gorunning
endif
DIR_WITH_MAIN=cmd/gorunning
TESTS        ?= ./...

build:
	cd ${DIR_WITH_MAIN} ;go build -o ${BINARY_NAME}

run:
	${BINARY_NAME}

test:
	set -a; go test $(TESTS)

inttest:
	set -a; go test $(TESTS) -tags=integration
