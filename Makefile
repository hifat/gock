run:
	go run ./cmd/api/main.go

migrate:
	go run ./cmd/migrate/main.go

test-repo:
	go test ./internal/repository/... -v