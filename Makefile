run:
	go run cmd/app/main.go

migration_up: 
	migrate -path ./migration/ -database "postgresql://postgres:postgres@localhost:5432/messanger?sslmode=disable" -verbose up

migration_down: 
	migrate -path ./migration/ -database "postgresql://postgres:postgres@localhost:5432/messanger?sslmode=disable" -verbose down

migration_fix:
	migrate -path ./migration/ -database "postgresql://postgres:postgres@localhost:5432/messanger?sslmode=disable" force 1