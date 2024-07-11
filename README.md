# Mini Online Shop Backend

This repository contains the backend for a mini online shop website. It handles user authentication, product management, shopping cart functionality, order processing, and payment integration.

## Features

- **User Authentication**: Login and registration functionality.
- **Product Management**: List and manage products.
- **Shopping Cart**: Add, update, and remove items in the cart.
- **Order Processing**: Place and manage orders (only buyer).
- **Payment Integration**: Handle payment processing (dummy).

## Tech Stack

- **Language**: Go
- **Framework**: Fiber
- **Database**: PostgreSQL, Redis

## Frameworks and Libraries

- **[GoFiber (HTTP Framework)](https://github.com/gofiber/fiber)**: An Express-inspired web framework for Go.
- **[GORM (ORM)](https://github.com/go-gorm/gorm)**: The fantastic ORM library for Go.
- **[Viper (Configuration)](https://github.com/spf13/viper)**: A complete configuration solution for Go applications.
- **[Golang Migrate (Database Migration)](https://github.com/golang-migrate/migrate)**: Database migrations for Golang.
- **[Go Playground Validator (Validation)](https://github.com/go-playground/validator)**: A Go library for struct and field validation.

## Entity-Relationship Diagram

Here is the ERD for the database schema:

![ERD](./docs/erd.png)

## Project Structure

```sh
.
├── cmd # Entry point of the application
│   └── core
├── db # Database migrations
│   └── migrations
└── internal
    ├── config # Application configs
    ├── delivery
    │   └── http
    │       ├── controller # Route handlers
    │       ├── middleware # Custom middlewares
    │       └── route # Application routes
    ├── domain # Delivery models
    ├── model # Database models
    ├── repository # Database logic
    ├── usecase # Business logic
    └── utils # Utility functions and helpers
```

## Getting Started

### Running locally

#### Prerequisites

- Go 1.20+
- PostgreSQL
- Redis

#### Set up the database

Connect to your PostgreSQL server, create the database, and do migration.

```sh
# Create the database
createdb db_mini_online_shop

# Apply migrations
migrate -path db/migrations -database postgres://[username]:[password]@localhost:5432/db_mini_online_shop?sslmode=disable up
```

#### Installation

Clone the repository:

```sh
git clone git@github.com:pravastacaraka/go-mini-online-shop-ws.git

cd go-mini-online-shop-ws
```

Install dependencies:

```sh
go mod vendor && go mod tidy
```

#### Configure the application

1. Open `config.local.json` file to configure the app config.
2. Change the database dan redis credentials according to your previous setting via environment variables look at the `.env.example` file.

#### Run the application

```sh
go run cmd/core/main.go
```

### Running container

Having trouble running manually in local? You can directly use the docker that I have provided.

#### Prerequisites

- Docker

#### Installation

```sh
# build the image
docker compose build

# start the app
docker compose up

# stop the app gracefully
docker compose down
```

## API Documentation

You can see the API documentation on here [Postman Collection](https://www.postman.com/avionics-physicist-27879440/workspace/mini-online-shop).
