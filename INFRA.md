# Development Infrastructure

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

## TODOs
* Logging Strategy
* Port allocation

## MongoDB

### First Issue - How to add data to MongoDB development container?
The service in a container method (`docker-compose`), plus docker principles enforce immutability. Data Volumes must also take this approach:

### Solution
* Create a separate node.js service (development-only) to take schema definition data (MongoDB JSON validation) and generate mock MongoJSON data
* Link mongoDB container and once MongoDB service starts use `mongoimport` to import data
