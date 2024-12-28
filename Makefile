initial-setup:
	@make build
	@make up 
setup:
	@make build
	@make up 
build:
	docker-compose build --no-cache --force-rm
stop:
	docker-compose stop
up:
	docker-compose up -d
compile:
	@make stop
	@make build
	@make up
