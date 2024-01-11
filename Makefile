generate-migrate:
	go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./migrations

go-osn-bash:
	go run cmd/sso/main.go --config=C:/Users/Алёна Валерьевна/Desktop/sso/сonfig/local.yaml