LINE1 ?= $(USER)

default: all

all: build run compile

build:
	go build

run:
	./goodnotes-cover --line1 "$(LINE1)"

compile:
	@script/compile.sh
