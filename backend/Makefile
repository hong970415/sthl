schemabuild:
	@echo "ent generate..."
	go generate ./ent
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert ./ent/schema

tidy:
	@echo "go mod tidy..."
	go mod tidy

docsbuild:
	@echo "updating swagger..."
	swag init

test_auth:
	@echo "Unit Testing auth package..."
	go test -race -count=1 -v ./authentication/...

test_dto:
	@echo "Unit Testing dto package..."
	go test -race -count=1 -v ./dto/...

test_service:
	@echo "Unit Testing service package..."
	go test -count=1 -v ./service/...

test_api:
	@echo "Unit Testing api package..."
	go test -count=1 -v ./api/... -short

test_api_integration:
	@echo "Integration Testing api package..."
	go test -count=1 -v ./api/...

test_all:
	@echo "Unit Testing all package..."
	go test -race -count=1 -v ./... -short