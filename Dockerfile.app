FROM golang:1.22

# Install PostgreSQL client
# to working with wait-for-db.sh
RUN apt-get update && apt-get install -y postgresql-client && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/src/app

COPY . .

RUN go mod download && go mod verify
RUN go build -v -o /usr/local/bin/app ./cmd/core

# Use wait-for-db script to wait for the database
CMD ["sh", "-c", "./wait-for-db.sh postgres -- app"]
