go mod tidy


GO_ENV=dev go run migrate/migrate.go
GO_ENV=dev go run main.go
