all: build

build: .
	go build -o built/jls .
	