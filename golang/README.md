go mod tidy


GO_ENV=dev go run migrate/migrate.go
GO_ENV=dev go run main.go


テストをするために、
address渡しをしてから実態を変更していたものを
値をreturnするように変更したが、これが良いのかどうかがわからない。

