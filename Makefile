# meta
NAME := SAMPLE_PRO
VERSION := $(shell git describe --abbrev=0 --tags)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.revision=$(REVISION)'

# set up the necessary tools
## SETUP
setup:
    go get github.com/golang/lint