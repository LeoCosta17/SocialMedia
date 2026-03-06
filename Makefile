# Variáveis para não repetir código
DB_URL=postgres://postgres:postgres@123!*@localhost/social?sslmode=disable
MIGRATE_PATH=./cmd/migrate/migrations

# Atalho para rodar as migrações pra cima
migrate-up:
	migrate -path=$(MIGRATE_PATH) -database="$(DB_URL)" up

# Atalho para forçar uma versão (ex: make migrate-force v=1)
migrate-force:
	migrate -path=$(MIGRATE_PATH) -database="$(DB_URL)" force $(v)