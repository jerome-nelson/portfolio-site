## Source code behind Portfolio Website

## How to Run
* `docker-compose up -d` ()
* `docker-compose config` (show printed configuration)

### Purpose of this Website
* To show technological experience
    - So aside from base technologies will attempt to minimise usage of libraries
* Improve knowledge of API building
* Get used to Go-Lang

### Technologies
* VSCode - for development
* MongoDB (for storing data)
* Go-Lang for API interface
* ReactJS for Website
* Docker/Docker-Compose for Development

## TODO
* Why each technology was picked
* MongoDB security best practices
    - Create sub-db with user roles

## Docker Services (for Development)
I use an .env file which `docker-compose`

| Service   | ENV PORT | Dev. Port Defaults  |   |   |
|-----------|----------|---------------------|---|---|
| MongoDB   | DB_PORTS | 27017 - 27019       |   |   |
| Go        | API_PORT | 1080                |   |   |

# Go Service
* `-port=<int>` launch main.go with this argument to start

# Reference Points
* https://docs.mongodb.com/manual/reference/bson-types/