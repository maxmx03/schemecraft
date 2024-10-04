clean:
	rm -rf ./build
build: clean
	go build .
	./schemecraft solarized.yml solarized-light.yml
