#!/bin/sh

PROGNAME=mongo_clean_sessions
PLATFORMS="darwin/386 darwin/amd64 \
dragonfly/amd64 \
freebsd/386 freebsd/amd64 freebsd/arm \
linux/386 linux/amd64 linux/arm linux/arm64 linux/ppc64 linux/ppc64le \
netbsd/386 netbsd/amd64 netbsd/arm \
openbsd/386 openbsd/amd64 openbsd/arm \
solaris/amd64 \
windows/386 windows/amd64"

VERSION=$(git tag -l | sort | tail -n1)
BUILD_TIME=`date +%FT%T%z`

git push --all
git push --tags
sleep 1
go get -u gopkg.in/mgo.v2

for PLATFORM in ${PLATFORMS}; do
    make clean > /dev/null 2>&1
    OS=${PLATFORM%/*}
    ARCH=${PLATFORM#*/}
    GOOS=${OS} CGO_ENABLED=0 GOARCH=${ARCH} go build -ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME}" -o ${PROGNAME} main.go
    ARCHIVE=${PROGNAME}-${VERSION}-${OS}-${ARCH}.tar.gz
    tar -czf ${ARCHIVE} ${PROGNAME}
    echo ${ARCHIVE}
done

make clean > /dev/null 2>&1
