clean:
	rm -rf ./build
build: clean
	go build .
	./schemecraft ./example/base46.yml
