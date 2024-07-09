# Stage 1: Build the application
FROM golang:1.22 as builder

ENV APP_HOME /go/src/app

WORKDIR ${APP_HOME}

COPY cmd/ cmd/
COPY internal/ internal/
COPY go.mod .
COPY go.sum .

RUN go mod download && go mod verify
RUN go build -o app ./cmd/core

# Stage 2: Runtime image
FROM golang:1.22

ENV APP_HOME /go/src/app

WORKDIR ${APP_HOME}

COPY db/ db/
COPY config.*.json .
COPY entrypoint.sh .
COPY --from=builder ${APP_HOME}/app .

# Install PostgreSQL client to check the PostgreSQL connection
RUN apt-get update && apt-get install -y postgresql-client && rm -rf /var/lib/apt/lists/*

# Install golang-migrate to do db migration
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN chmod +x ./entrypoint.sh

EXPOSE ${PORT}

ENTRYPOINT ["./entrypoint.sh"]
CMD ["./app"]
