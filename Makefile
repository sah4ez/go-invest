SOURCE=$(shell find ./* -type f | grep .*go)
STACK_NAME=INV
BINARY_NAME=goinvest

.PHONY: run
run:
	@env $(shell cat .env | xargs) ./$(BINARY_NAME)

build: $(SOURCE) vendor
	go build -o $(BINARY_NAME)

.PHONY: test
test: $(SOURCE)
	go test $(shell glide novendor)

.PHONY: migration
migration: migrations vendor
	@env $(shell cat .env | xargs) go run main.go --migration

.PHONY: prepare
up: docker-compose.yml
	docker stack deploy --compose-file docker-compose.yml $(STACK_NAME)

.PHONY: vendor
vendor: glide.lock glide.yaml glide_up.sh
	$(shell ./glide_up.sh)

.PHONY: down
down:
	docker stack rm $(STACK_NAME)
