# go-passwordvalidator

The service was designed on a layer oriented architecture with influences on the\
[Ports and Adapters of Hexagonal Architecture](https://dev.to/jofisaes/hexagonal-architecture-ports-and-adapters-1h4m).

It consist basically on some very well defined layers:

- **Application** holds the application configuration (e.g. env vars) and its bootstrap (servers, queue consumers, etc).
- **Domain** holds everything related to business logic, such entities and services.
- **Resources** consists on everything related to communication with external resources, such APIs, Repositories, etc. This service \
does not need this layer, so it's not defined.

In a more robust scenario, all layers should make heavy use of interfaces, but given the little scope of this service, \
only the business layer (domain/services) is using an interface in order to be abstracted and consequently decoupled.

So basically, we have the application being started and under it also is an http server that serves on the port `7000` \
and the route `/validate`, with requests being handled by the controller, which calls the service interface in order \
have all the business logics applied to the request data that finally retrieves the desired information back to the \
controller and so on.

- Application -> HTTP Server -> Controller -> Service

The server also is configured to handle shutdown signals in a graceful away, so all the ongoing connections can have \
time to be finished and give response to the client before the server being shutdown.

## IsValid

The main function responsible for the password validation uses other private functions that represent each of the \
validation steps: the size and the missing required characters. It keeps a `map` of each Rule used to define some \
criteria, so, this way, only one iteration can be done on the entire string. If the size does not match or any invalid \
character is found, the iteration does not happen or stops immediatelly.

## Usage

Run with Docker Compose

```shell script
docker-compose up --build -d
```

Or with `make`, just run
```shell script
make run
```

**ATTENTION:** if you want to edit any configuration available (such as port, timeouts and password sizes), please, do it in the `docker-compose` file.

The server will expose an HTTP API at port 7000 on `localhost`.

Endpoint `/validate`
Method &nbsp;&nbsp;`POST`
Request body example
```json
{
  "password": "blaFOOabcdefg"
}
```
Response body example
```json
{
  "isValid": false,
  "errors": [
    "should have at least one digit",
    "should have at least one of the special characters: !@#$%^&*()-+"
  ]
}
```
Status Code: the API will always return `200` for succeeded requests, even for invalid passwords. But it can also return `400` for bad request.

## Testing

Unit tests are made using the pattern of table tests, which facilitates the reading of inputs and expected outputs based on each scenario.

To run unit tests (including race condition), execute
```shell script
make test/unit
```

To run integration tests (including race condition), execute
```shell script
make test/integration
```

Or run both at once with
```shell script
make test/all
```
