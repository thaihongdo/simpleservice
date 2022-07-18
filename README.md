
  

# Golang Simple Api

  

A Golang REST API built with Go Gin, connected to SQLite Database. The main functions:
  
Api request for wager:  
  

## System

  

---

  

This repository was developed with the system versions below:

  

- Go 1.18

- MacOS (or Ubuntu 18.x)

  

## How to use

---
### Run from source code

Clone this repository to your local machine. Ensure that you have Go environment. Go to root project and run:
```
make run 
OR
go run cmd/entity-server/main.go
```
### Run by docker

From root project run the script:
```
bash ./scripts/run.dev.sh
```


Sample request:
```
curl --location --request POST 'http://localhost:8080/wagers' \
--header 'Content-Type: application/json' \
--data-raw '{
    "total_wager_value": 1,
    "odds": 1,
    "selling_percentage": 1,
    "selling_price": 3
}'
```

