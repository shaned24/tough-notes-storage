IMAGE=shaned24/tough-notes
TAG=latest
DEV_IMAGE=shaned24/tough-notes-dev
DEV_TAG=latest
DC=docker-compose
DC_DEV_FLAGS=-f docker/compose/base.yaml
DEV_ENV=$(DC) $(DC_DEV_FLAGS)
GO_PACKAGE=github.com/shaned24/tough-notes-storage
HELM_CHART=deploy/helm/toughnotes
HELM_RELEASE_NAME=toughnotes


#TEST_COMMAND=go test ./... -cover -coverpkg ./... -coverprofile=c.out
TEST_COMMAND=go test ./...

.PHONY: deploy

generate:
	@protoc api/notespb/notes.proto --go_out=plugins=grpc:.

docker-build:
	@docker build . -f docker/Dockerfile -t $(IMAGE):$(TAG)

docker-build-dev:
	@docker build --target dev . -f docker/Dockerfile -t $(DEV_IMAGE):$(DEV_TAG)

docker-test: docker-build-dev
	@docker run --rm $(DEV_IMAGE):$(DEV_TAG) $(TEST_COMMAND)

docker-publish: docker-build
	@docker push $(IMAGE):$(TAG)

test:
	$(TEST_COMMAND)

env:
	$(DEV_ENV) up -d

logs:
	$(DEV_ENV) logs -f

run:
	go run .

client-run:
	go run notes/client/client.go

docker-run:
	@docker run --rm $(IMAGE):$(TAG)

mocks:
	mockgen \
	 -package tests \
	 -destination internal/pkg/storage/tests/mocks.go $(GO_PACKAGE)/internal/pkg/storage NoteStorage

clean:
	rm c.out
	rm coverage.html

coverage:
	go tool cover -html=c.out -o coverage.html

deploy:
	HELM_RELEASE_NAME=$(HELM_RELEASE_NAME) \
	HELM_CHART=$(HELM_CHART) \
	./scripts/deploy.sh

wire:
	wire