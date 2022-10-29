## Balance App

This document and openapi file are still in "work in progress" state.

## Setup

All the following instructions are for Unix based systems.
For Windows, commands might slightly diverse.

### Requirements

To start the app as the docker container you need to have `docker` and `docker-compose` installed.

- I used `docker-compose` version `v2.12.0` and `docker` version `20.10.20`.

To run the container execute the following command:

```bash:
docker compose up
```

To start the app on your local machine you need to have `Golang` and `PostgreSQL` installed.

- I used `Golang` version `1.19.1` and `PostgreSQL` version `13.4`.

To run app on the local device firstly make sure that **local.env** file contains accurate credentials of your DB, than –– execute the following command:

```bash:
go run ./cmd/api/main.go
```

## API

User

`GET http://host/users` - get all users  
`GET http://host/users/:id` - get user info by id  
`GET http://host/users/:id/balance` - get user balance by id

Transaction

`POST http://host/transactions/deposit` - deposit money to user balance  
`POST http://host/transactions/freeze` - freeze money for the order within a service on user balance  
`POST http://host/transactions/apply` - apply frozen transaction  
`POST http://host/transactions/revert` - reverts frozen transaction  
`POST http://host/transactions/stat/:id` - get user transactions  
`POST http://host/transactions/services-sum-amount` - get link to csv, containing total money spent on services within month and year
