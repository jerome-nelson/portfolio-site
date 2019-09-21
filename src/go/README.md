## Go-Lang API Service

This service is responsible for connecting to MongoDB instance and exposing data via HTTP endpoints. Responses are returned as JSON (since API is for React Application)

* `port` - The port to run the server on (required)

### TODO
* Re-examine Regex for API
* Make api getData monadic
* Allow API communication under specific IPs only
* Error Mapping To map all common expected errors
* COnvert arrays to result

## Contract
* Any calls to `/api/` with incorrect url, return 404
* All errors are added to `ErrorMsg` as part of payload response
* Reponse is empty if

## Error Types
* JSON Error
* Incorrect Service
* Regular 404

#### Cases to cover
1. Handling error (if service doesn't exist)
2. if port is busy throw exception

## TO INVESTIGATE
* Async Messaging
* Scaling (Replication)
* Metrics (Prometheus Metric)
* Caching (ETCD)
* Bee's with Machine Games (Testing Scaling)

## External References
1. [https://godoc.org/go.mongodb.org/mongo-driver/mongo](Go MongoDB API)