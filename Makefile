IMAGE=shaned24/tough-notes
DEV_IMAGE=shaned24/tough-notes-dev
DC=docker-compose
DC_DEV_FLAGS=-f docker/compose/base.yaml
DEV_ENV=$(DC) $(DC_DEV_FLAGS)
GO_PACKAGE=github.com/shaned24/tough-notes-storage

generate:
	@protoc notes/notespb/notes.proto --go_out=plugins=grpc:.

build:
	@docker build . -f docker/Dockerfile -t $(IMAGE):latest

build-dev:
	@docker build . -f docker/Dockerfile.dev -t $(DEV_IMAGE):latest

test:
	@docker run --rm -v $(PWD):/workspace $(DEV_IMAGE):latest go test ./...

env:
	$(DEV_ENV) up -d

logs:
	$(DEV_ENV) logs -f

run:
	go run notes/server/server.go

client-run:
	go run notes/client/client.go

docker-run:
	@docker run --rm $(IMAGE):latest

mocks:
	mockgen \
	 -package tests \
	 -destination notes/server/storage/tests/mock_storage.go $(GO_PACKAGE)/notes/server/storage NoteStorage
