new:
	dotnet new $(type) -o $(name)

add:
	dotnet sln $(sln).slnx add $(name)/$(name).csproj

run:
	dotnet run --project $(name)/$(name).csproj

build:
	dotnet build $(name).slnx

sln:
	dotnet new sln -n $(sln)

# type choices: webapi (API), worker (Background), console (CLI), classlib (Library), xunit (Test)
