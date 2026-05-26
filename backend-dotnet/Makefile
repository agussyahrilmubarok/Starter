SLN ?= BackendApi
# TYPE choices: webapi (API), worker (Background), console (CLI), classlib (Library), xunit (Test)
TYPE ?= webapi

.PHONY: setup create/sln create/project add/project run build clean

setup:
	@echo "==> Membuat Solution: $(SLN).sln..."
	dotnet new sln -n $(SLN)
	@echo "==> Membuat Project $(TYPE) di folder $(NAME)..."
	dotnet new $(TYPE) -o $(NAME)
	@echo "==> Menghubungkan Project ke Solution..."
	dotnet sln $(NAME).slnx add $(NAME)/$(NAME).csproj
	@echo "==> Setup selesai!"

create/sln:
	dotnet new sln -n $(SLN)

create/project:
	dotnet new $(TYPE) -o $(NAME)

add/project:
	dotnet sln $(SLN).slnx add $(NAME)/$(NAME).csproj

run:
	dotnet run --project $(NAME)/$(NAME).csproj

build:
	dotnet build $(NAME).sln

clean:
	dotnet clean $(NAME).sln
	rm -rf $(NAME)/obj $(NAME)/bin