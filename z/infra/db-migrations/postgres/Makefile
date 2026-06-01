.PHONY: migrate migrate-info migrate-validate migrate-repair migrate-clean

FLYWAY = docker run --rm \
	--add-host=host.docker.internal:host-gateway \
	-v $(PWD)/sql:/flyway/sql \
	-v $(PWD)/flyway.conf:/flyway/conf/flyway.conf \
	flyway/flyway:10.21.0

migrate:
	$(FLYWAY) migrate

migrate-info:
	$(FLYWAY) info

migrate-validate:
	$(FLYWAY) validate

migrate-repair:
	$(FLYWAY) repair

migrate-clean:
	$(FLYWAY) clean