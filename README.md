# Shippy

Go lang microservices project for shipping containers

## Prerequisits

## Development
Use the make file for common tasks

1. Build the proto files `make proto`
2. Build the service `make build`
3. Build the Docker container `make docker`

### Micro web client
This can be use to test the microservices

Start the service you want to test, then also boot the micro web service using the `micro` cli.
```
$> micro cli
```
View in browser at `http://127.0.0.1:8082/client`

## Deployment
