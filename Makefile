.PHONY: *

run:
	-make down
	docker-compose -f ./deploy/docker-compose.yml up --build

down:
	docker-compose -f ./deploy/docker-compose.yml down --remove-orphans
