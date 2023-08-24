run:
	go run ./cmd/api/main.go

migrate:
	go run ./cmd/migrate/main.go

mock-gen:
	go generate ./internal/service/...;
	go generate ./internal/repository/...;

test-service:
	go test ./internal/service/... -v

test-repo:
	go test ./internal/repository/... -v

test:
	go test ./...