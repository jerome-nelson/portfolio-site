# Portfolio Website
This repository is the source behind my website [link here](https://www.test.test). It's a SoC/Microservice type stack
using different languages to store/process and display data in a React UI.

## How to Install

**Note: runs on Docker Linux Containers only**

Install [Docker](https://www.docker.com/products/docker-desktop). Then create `.env` file in root with similar values to below:
```
API_PORT=<PORT>
DB_PORT=27017
DB_ROOT_USER=root
DB_ROOT_PASS=pass
DB_USER=site
DB_PASS=pass
DB_HOST=localhost
DB_NAME=site_db
```

After that run `docker-compose up`

## Short-term Aims
* Knowledge building
* Improve knowledge of API building
* Increase familiarity with Go-Lang


## Long-term Aims
* Deploy to AWS Cluster
* Understanding of Production/Development Infrastructure Lifecycles
* Load-testing/Metric's Monitoring (This isn't needed but the experience is)

### Technologies

| Tech.         | Responsibility                                                       |                        |
|---------------|----------------------------------------------------------------------|------------------------|
| MongoDB       | Store long term data persistently                                    | [Docs](docs/MONGO.md)  |
| React         | Thin-client responsible for rendering data                           |                        |
| Go-Lang       | API layer responsible for aggregation/formattting of data            | [Docs](docs/GO.md)     |
| Docker        | Used as a service mesh to allow independent linked services          | [Docs](docs/DOCKER.md) |


### Useful commands
* `docker-compose up`
* `docker-compose config` (show printed configuration)


## External Tools
* [Wait-for-it](https://github.com/vishnubob/wait-for-it) - Shell script which waits for available connection before
running a command