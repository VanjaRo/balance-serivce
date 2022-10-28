Balance App

This document and openapi file are still in "work in progress" state.

To start the serice run:
`docker compose up`

All api handlers available:

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
