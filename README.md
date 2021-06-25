
# Tinier Auth Service
Tinier is a URL Shortener REST API written in go.

## Installation
Open docker-compose.yml and set env variables for cassandra and redis addresses properly.
```
$ git clone https://github.com/pooladkhay/tinier-auth-service.git
$ cd tinier-auth-service
$ docker-compose up --build
```
## Design

![Ttinier](tinier-design.png?raw=true)