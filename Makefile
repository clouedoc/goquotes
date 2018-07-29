all: goquotes

goquotes: main.go
	go build -o goquotes main.go

test: goquotes
	./goquotes programming
