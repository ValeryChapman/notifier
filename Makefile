start-dev:
	docker-compose -f docker-compose.dev.yml up --build
start-dev-d:
	docker-compose -f docker-compose.dev.yml -d up --build
start-prod:
	docker-compose -f docker-compose.prod.yml up --build
start-prod-d:
	docker-compose -f docker-compose.prod.yml up -d --build
