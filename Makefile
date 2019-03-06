IMAGE=shaned24/tough-notes
DEV_IMAGE=shaned24/tough-notes-dev
DC=docker-compose
DC_DEV_FLAGS=-f docker/compose/base.yaml
DEV_ENV=$(DC) $(DC_DEV_FLAGS)
GO_PACKAGE=github.com/shaned24/tough-notes-storage
HELM_CHART=deploy/helm/toughnotes
HELM_RELEASE_NAME=toughnotes

.PHONY: deploy

generate:
	@protoc notes/notespb/notes.proto --go_out=plugins=grpc:.

build:
	@docker build . -f docker/Dockerfile -t $(IMAGE):latest

build-dev:
	@docker build . -f docker/Dockerfile.dev -t $(DEV_IMAGE):latest

test:
	go test ./... -cover -coverpkg ./... -coverprofile=c.out

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

clean:
	rm c.out
	rm coverage.html

coverage:
	go tool cover -html=c.out -o coverage.html

publish-image: build
	@docker push $(IMAGE)

deploy:
	@helm upgrade \
	 --install \
	 $(HELM_RELEASE_NAME) \
	 $(HELM_CHART)