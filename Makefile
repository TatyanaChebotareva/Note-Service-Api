PHONY: generate
generate:
	mkdir -p pkg/note_v1

	protoc 	--proto_path api/note_v1	\
			--go_out=pkg/note_v1 --go_opt=paths=source_relative \
			--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=source_relative \
			api/note_v1/note.proto


LOCAL_MIGRATIONS_DIR=./migrations
LOCAL_MIGRATIONS_DSN="host=localhost port=54321 dbname=note-service user=note-service-user password=note-service-password"

.PHONY: install-goose
.install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY:	local-migrations-status
local-migrations-status:
	goose -dir ${LOCAL_MIGRATIONS_DIR} postgres ${LOCAL_MIGRATIONS_DSN} status -v

.PHONY:	local-migrations-up
local-migrations-up:
	goose -dir ${LOCAL_MIGRATIONS_DIR} postgres ${LOCAL_MIGRATIONS_DSN} up -v

.PHONY:	local-migrations-down
local-migrations-down:
	goose -dir ${LOCAL_MIGRATIONS_DIR} postgres ${LOCAL_MIGRATIONS_DSN} down -v