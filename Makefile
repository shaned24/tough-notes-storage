DC=docker-compose
DC_DEV_FLAGS=-f docker/compose/base.yaml

DEV_ENV=$(DC) $(DC_DEV_FLAGS)

generate:
	@protoc notes/notespb/notes.proto --go_out=plugins=grpc:.

env:
	$(DEV_ENV) up -d

logs:
	$(DEV_ENV) logs -f

run:
	go run notes/server/server.go