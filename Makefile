
PHONY: install
install:
	go install github.com/swaggo/swag/cmd/swag@v1.8.12
	go install github.com/google/wire/cmd/wire@v0.5.0
	go install github.com/cosmtrek/air@v1.42.0

PHONY: demo
demo:
	docker compose -f ./docker/compose.yaml up

PHONY: stop
stop:
	docker compose -f ./docker/compose.yaml stop

PHONY: clean
clean:
	docker container rm db-mysql app

PHONY: develop
develop:
	docker compose -f ./docker/compose.yaml run --service-ports --name dev-mysql -d mysql
	air