SLN ?= BackendApi
# TYPE choices: webapi (API), worker (Background), console (CLI), classlib (Library), xunit (Test)
TYPE ?= webapi

.PHONY: setup create/sln create/project add/project run build clean

setup:
	dotnet new sln -n $(SLN)
	dotnet new $(TYPE) -o $(NAME)
	dotnet sln $(NAME).slnx add $(NAME)/$(NAME).csproj

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