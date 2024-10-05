clean:
	rm -rf ./build
build: clean
	go build .
	./schemecraft dracula.yml dracula-soft.yml
