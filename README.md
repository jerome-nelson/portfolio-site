# Portfolio Website
This repository is the soruce behind my website [link here](https://www.test.test). It's a SoC/Microservice type stack
using different languages to store/process and display data in a React UI.

## Why?
* To show technological experience
    - So aside from base technologies will attempt to minimise usage of libraries
* Improve knowledge of API building
* Become familiar with Go-Lang

### Technologies

| Tech.         | Responsibility                                                       |
|---------------|----------------------------------------------------------------------|
| MongoDB       | Store long term data persitently                                     |
| React         | Thin-client responsible for rendering data                           |
| Go-Lang       | API layer responsible for aggregation/formattting of data            |
| Docker        | Used as a service mesh to allow independent linked services          |

To learn more please take at the source code/or the README.md in each `src/<folder>`


### How to Run
* `docker-compose up -d` ()
* `docker-compose config` (show printed configuration)

## TODO
* Combine Volumes
* Create network (if needed)
* Fix Dev MongoSeeding
    - Auto import of mongoDB
* Global Configuration
    - For connecting DB collections with GO API
* MongoDB security best practices
    - Create website DB with user roles (done)
* MongoDB Validation

## External Tools
* (https://github.com/vishnubob/wait-for-it)[Wait-for-it] - Shell script used to SEED MongoDev Container