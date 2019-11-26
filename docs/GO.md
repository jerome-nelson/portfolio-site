## Go-Lang API Service

This simple service is responsible for connecting to MongoDB instance and exposing data.

### Responses
* All Non-GET requests return `405 - Method not allowed`
* Successful Responses are returned as JSON with `200`
* Critical Errors cause application to exit
    - DB Connection errors
    - Application Startup errors
* Errors will throw `500` with `Error Occurred` displayed
    - Actual error will be logged on system


### Application Flags (all required!)
* `serve`
 - The port to run the server on (required)
* `dburl`
 - Full mongoDB schema url `mongodb://<user>:<pass>@<host>:<port>`
* `dbname`
 - Name of MongoDB DB to query

## TO INVESTIGATE
* Gorilla Mux (compare solution ~ possibly use for React SSR)
* Async Messaging
* Scaling (Replication)
* Metrics (Prometheus Metric)
* Caching (ETCD)
* Bee's with Machine Games (Testing Scaling)

## External References
1. [https://godoc.org/go.mongodb.org/mongo-driver/mongo](Go MongoDB API)
2. [https://blog.golang.org/defer-panic-and-recover](Defer, Panic, Recover)
3. [https://www.alexedwards.net/blog/organising-database-access](DB Persistence - handling connections)

### TODOS
* Gather all requirements for rewrite
* Re-examine Regex for API
* Make api getData monadic
* Allow API communication under specific IPs only
* Error Mapping To map all common expected errors
* COnvert arrays to result
