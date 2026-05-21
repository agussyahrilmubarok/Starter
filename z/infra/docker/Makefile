compose/up:
	docker compose -f docker-compose-$(NAME).yml up -d --build

compose/down:
	docker compose -f docker-compose-$(NAME).yml down
	docker volume prune -a
