DB_URL=mysql://user:password@tcp(localhost:3306)/db


migrate-up:
	migrate -path ./internal/database/migrations -database "${DB_URL}" up

migrate-down:
	migrate -path ./internal/database/migrations -database "${DB_URL}" down

run:
	go run cmd/main.go

