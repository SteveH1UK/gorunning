# A Practice Go application (MongoDb and Rest API)
This application allows an  Web Client to call the rest api to populate a Mongo DB Record. Tested on ubuntu.

Database collections:
* Athlete
    * Id
    * Name
    * Friendly Name
    * Date of Birth
    * Creation Date


# Rest interface 
All are preceded by context (which ends in gorunning)

| Verb | URL | Description | Notes
| ---- | ------------ | ---------- | ---
| POST | atheletes | Add an athelete | Returns 201 on success
| GET  | atheletes | Lists all atheletes |
| GET  | athelete/{friendly_name} | List one athelete |
| PUT  | athelete/{friendly_name} | Update athelete  | Returns 200 on success and the updated record

# Package Structure
[Ben Johnsons Package Struture](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)

# References
* [Go Microservies with MongoDB](http://goinbigdata.com/how-to-build-microservice-with-mongodb-in-golang/)
* [How I write Go HTTP Servers (after 7 years)](https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831)


# Building
make build

# Set up and configuration
## Setup
### Infrastructure
* Golang (tested 1.10.3)
* MongoDB (tested on 3.6.3)

### External Packages
```
go get github.com/gorilla/pat
go get github.com/justinas/alice
go get github.com/ian-kent/gofigure
go get github.com/globalsign/mgo
```

## Configuration

| Environment variable | Flag | Description | Default Value
| -------------------- | ---- | ----------- | --------------
| `RUNNING_MONGO_URL` | db-name | MongoDb URL | localhost:27017
| `RUNNING_DB_NAME` | mongo-url | MongoDb Database Name | gorunning
| `RUNNING_LISTEN_ADDRESS` | listen-address | HTTP Server Listen Address | 127.0.0.1:8080
| `TEST_RUNNING_MONGO_URL` | test-db-name | Test MongoDb URL | localhost:27017
| `TEST_RUNNING_DB_NAME` | test-mongo-url | Test MongoDb Database Name | gorunning


# Running
On Command line:
make run

# Unit Testing
make test

# Unit and Integration Testing
make inttest

# System Testing using curl
## Insert record (Success = 201)
```
curl -d '{"friendly-url":"johndoe", "name":"John Doe", "date-of-birth":"18-10-1990"}' -H "Content-Type: application/json"  -X POST http://127.0.0.1:8080/gorunning/atheletes
curl -d '{"friendly-url":"sebcoe", "name":"Seb Coe", "date-of-birth":"11-10-1888"}' -H "Content-Type: application/json"  -X POST http://127.0.0.1:8080/gorunning/atheletes
curl -d '{"friendly-url":"johndoe16", "name":"John Doe16", "date-of-birth":"18-09-1990"}' -H "Content-Type: application/json"  -X POST  http:///127.0.0.1:8080/gorunning/atheletes
```
## Insert record (Validation Error, 422)
```
curl -d '{"friendly-url":"joe", "name":"John Doe", "date-of-birth":"1990-10-18"}' -H "Content-Type: application/json"  -X POST http://127.0.0.1:8080/gorunning/atheletes
```
## Get record (Success = 200)
```
curl  http://127.0.0.1:8080/gorunning/athelete/johndoe
```
## Get all records
```
curl  http://127.0.0.1:8080/gorunning/atheletes
```
## Edit
```
curl -d '{"friendly-url":"johndoe", "name":"John Doe", "date-of-birth":"11-03-1986"}' -H "Content-Type: application/json"  -X PUT http://127.0.0.1:8080/gorunning/athelete/johndoe
curl -d '{"friendly-url":"JohnFDoe", "name":"John Doe", "date-of-birth":"11-03-1986"}' -H "Content-Type: application/json"  -X PUT http://127.0.0.1:8080/gorunning/athelete/johndoe16
```

# Testing with external dependencies
When using structs with methods you also need to create an interface. In this case you a framework to generate the mocks for you

## What to test with add_athelete.go
* JSON parsed OK to NewAthelete struct 
* JSON parse error
* Validation Errors (all)
* Error from DAO (both types)
* Happy Path - check return

## Using Gomock to test add_athelete.go
* Create mocks folder at root level
* mockgen -destination=mocks/mock_athelete_dao.go -package=mocks github.com/SteveH1UK/gorunning/mongodb AthleteDAOInterface

