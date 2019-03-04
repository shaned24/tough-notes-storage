FROM golang

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /tough-notes notes/server/server.go

FROM scratch

# Copy the artifact from the first build stage into the second stage (which will
# become the final image)
COPY --from=0 /tough-notes .

CMD ["./tough-notes"]