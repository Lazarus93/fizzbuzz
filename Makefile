.PHONY: build
build:
	go build -o main ./cmd/fizzbuzz

.PHONY: test
test:
	go test -v ./...

.PHONY: doc
doc:
	swag init

.PHONY: docker
docker:
	docker build -t fizzbuzz .

.PHONY: run
run:
	./main


.PHONY: clean
clean:
	rm main
