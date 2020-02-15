# Development
FROM golang:1.13 as dev

RUN go get github.com/google/wire/cmd/wire && \
    go get github.com/golang/mock/gomock && \
    go get github.com/golang/protobuf/protoc-gen-go

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .

# Builder
FROM dev as builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /tough-notes notes/server/main.go notes/server/wire_gen.go

# Production
FROM scratch
# Copy the artifact from the first build stage into the second stage (which will
# become the final image)
COPY --from=builder /tough-notes .
CMD ["./tough-notes"]