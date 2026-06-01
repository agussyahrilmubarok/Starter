.PHONY: migrate migrate-status

CQLSH = docker exec -it cassandra cqlsh -u cassandra -p cassandra

migrate:
	@for f in $(shell ls $(PWD)/cql/V*.cql | sort); do \
		echo "Running migration: $$f"; \
		docker exec -i cassandra cqlsh -u cassandra -p cassandra < $$f; \
	done

migrate-status:
	$(CQLSH) -e "SELECT * FROM system_schema.keyspaces;"