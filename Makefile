SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=mongo_clean_sessions

VERSION=`git tag -l | sort | tail -n1`
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME}"

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	go get gopkg.in/mgo.v2
	go build ${LDFLAGS} -o ${BINARY} main.go

.PHONY: install
install:
	go install ${LDFLAGS} ./...

.PHONY: fclean
fclean: clean
	rm -f *.tar.gz ~* .#*

.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm -f ${BINARY} ; fi

