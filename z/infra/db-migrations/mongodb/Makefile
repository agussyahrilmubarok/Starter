.PHONY: migrate migrate-info

MONGOSH = docker run --rm \
	--add-host=host.docker.internal:host-gateway \
	-v $(PWD)/js:/migrations \
	mongo:8.0 mongosh \
	"mongodb://sandbox:sandbox@host.docker.internal:27017/sandbox_db?authSource=admin"

migrate:
	@for f in $(shell ls $(PWD)/js/V*.js | sort); do \
		echo "Running migration: $$f"; \
		$(MONGOSH) /migrations/$$(basename $$f); \
	done

migrate-info:
	$(MONGOSH) --eval "db.getMigrationStatus()"