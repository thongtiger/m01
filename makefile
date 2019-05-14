include .env
export # export variable from read .env
env:
	@echo PATH=$(PATH)
	@echo PORT=$(PORT)
	@echo MSSQL_HOST=$(MSSQL_HOST)
	@echo MSSQL_DB=$(MSSQL_DB)
	@echo MSSQL_USERNAME=$(MSSQL_USERNAME)
	@echo MSSQL_PASSWORD=$(MSSQL_PASSWORD)
	@echo tsql_ExistUsername=$(tsql_ExistUsername)
	@echo tsql_GetProfile=$(tsql_GetProfile)
build_linux:
	env GOOS=linux GOARCH=amd64 go build
build_windows:
	env GOOS=windows GOARCH=amd64 go build
clean:
	go clean
run:
	go run .
test:
	go test .
docker:
	docker-compose -f "docker-compose.yml" up -d --build