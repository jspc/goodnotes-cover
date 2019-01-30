COVERNAME ?= $(USER)

default: all

all: build run

build:
	go build

run:
	./goodnotes-cover --name $(COVERNAME)
