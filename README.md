# Golang simple API project using Gin, Gorm, JWT and PostgreSQL 

In this project you are going to find:

* Basic routing
* JWT authentication 
* UserLogin and Register
* Custom middleware
* Migrations
* Project folder structure
* Basic Crud with GORM
* Docker environment for database

## Start
Go to /go-app folder and run

```bash
go run .
```

## Example of .env on the /go-app folder

PORT=":8888"
SECRET=auth-api-jwt-secret
GIN_MODE=debug
DB_HOST="192.168.1.122"
DB_PORT="5432"
DB_USER="postgres"
DB_PASSWORD="postgres"
DB_NAME="golang_test"

## First time? Run the migrations

Uncomment lines 5 and 30 of main.go

## Docker environment

Run _make initial-setup_ to deploy the docker setup

## Greeting

```
  while (true) {
    echo 'Thanks for watching and don't forget to rate this project!';
  }
```

Best,
Daniel