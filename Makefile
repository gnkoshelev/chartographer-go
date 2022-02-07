default: clean test build run

test:
	go test

build:
	go build -o build/app

run:
	./build/app

clean:
	rm -rf build