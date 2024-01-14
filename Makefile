generate-migrate:
	go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./migrations

go-osn-bash:
	go run cmd/sso/main.go --config=C:/Users/Алёна Валерьевна/Desktop/sso/сonfig/local.yaml
	go run cmd/sso/main.go --config=./сonfig/local.yaml

generate-tests-migrate:
	go run ./cmd/migrator/main.go --storage-path=./storage/sso.db --migrations-path=./tests/migrations --migrations-table=migrations_table

star-before-first-test:
	go run cmd/sso/main.go --config=./сonfig/local.yaml (suite.go cfg := config.MustLoadByPath("./сonfig/local.yaml"))

all-tests:
	go test ./tests