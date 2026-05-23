compose/up:
	docker compose -f docker-compose-$(name).yml up -d --build

compose/down:
	docker compose -f docker-compose-$(name).yml down
	docker volume prune -a
