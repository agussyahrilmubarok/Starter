PROJECT := BackendApi.csproj
CTX     ?= AppDbContext

.PHONY: setup/db dev pkg/add db/add db/update db/rollback db/list db/drop

setup/db:
	dotnet add package Npgsql.EntityFrameworkCore.PostgreSQL
	dotnet add package Microsoft.EntityFrameworkCore.Design
	dotnet tool install --global dotnet-ef || dotnet tool update --global dotnet-ef

dev:
	dotnet watch run --project $(PROJECT)

pkg/add:
	dotnet add package $(p)

db/add:
	dotnet ef migrations add $(m) --context $(CTX)

db/update:
	dotnet ef database update --context $(CTX)

db/rollback:
	dotnet ef migrations remove --context $(CTX)

db/list:
	dotnet ef migrations list --context $(CTX)

db/drop:
	@echo "==> WARNING: Akan menghapus database..."
	dotnet ef database drop --force --context $(CTX)