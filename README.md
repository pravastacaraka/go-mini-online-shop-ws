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

```bash
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

### Prerequisites

- Go 1.20+
- PostgreSQL
- Redis
- Docker

### Installation

Clone the repository:

```bash
git clone git@github.com:pravastacaraka/go-mini-online-shop-ws.git

cd go-mini-online-shop-ws
```

Install dependencies:

```bash
go mod vendor && go mod tidy
```

### Set up the database

Connect to your PostgreSQL server, create the database, and do migration.

```bash
# Create the database
createdb db-mini-online-shop

# Apply migrations
migrate -path ./migrations -database postgres://[username]:[password]@localhost:5432/db-mini-online-shop?sslmode=disable up
```

### Configure the application

1. Open `config.local.json` file in the root directory to configure the application config.
2. Change the database, redis credentials according to your previous setting via environment variables look at the `.env.example` file.

### Run the application

```bash
go run cmd/core/main.go
```

## API Documentation

You can see the API documentation on here [Postman Collection](https://www.postman.com/avionics-physicist-27879440/workspace/mini-online-shop).
