# Development Infrastructure

## TODOs
* Logging Strategy
* Port allocation

## MongoDB

### First Issue - How to add data to MongoDB development container?
The service in a container method (`docker-compose`), plus docker principles enforce immutability. Data Volumes must also take this approach:

### Solution
* Create a separate node.js service (development-only) to take schema definition data (MongoDB JSON validation) and generate mock MongoJSON data
* Link mongoDB container and once MongoDB service starts use `mongoimport` to import data
