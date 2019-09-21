# Portfolio Website
This repository is the source behind my website [link here](https://www.test.test). It's a SoC/Microservice type stack
using different languages to store/process and display data in a React UI.

## Short-term Aims
* To show technological experience
    - So aside from base technologies will attempt to minimise usage of libraries
* Improve knowledge of API building
* Increase familiarity with Go-Lang


## Long-term Aims
* Deploy to AWS Cluster
* Understanding of Production/Development Infrastructure Lifecycles
* Load-testing/Metric's Monitoring (This isn't needed but the experience is)

### Technologies

| Tech.         | Responsibility                                                       |
|---------------|----------------------------------------------------------------------|
| MongoDB       | Store long term data persistently                                     |
| React         | Thin-client responsible for rendering data                           |
| Go-Lang       | API layer responsible for aggregation/formattting of data            |
| Docker        | Used as a service mesh to allow independent linked services          |

To learn more please take at the source code/or the *.md files in `/docs`


### Useful commands
* `docker-compose up`
* `docker-compose config` (show printed configuration)


## External Tools
* [Wait-for-it](https://github.com/vishnubob/wait-for-it) - Shell script used to SEED MongoDev Container