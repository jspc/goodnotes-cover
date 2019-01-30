COVERNAME ?= $(USER)

default: all

all: build run compile

build:
	go build

run:
	./goodnotes-cover --name "$(COVERNAME)"

compile:
	@script/compile.sh
