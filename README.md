# Shippy

Go lang microservices project for shipping containers

## Prerequisits
Install:
1. MongoDB

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

Example JSON payloads that can be sent to services

#### Consignment JSON
```
{
  "description": "This is a test consignment",
  "weight": 55000,
  "containers": [
    { "customer_id": "cust001", "user_id": "user001", "origin": "Manchester, United Kingdom" },
    { "customer_id": "cust002", "user_id": "user001", "origin": "Derby, United Kingdom" },
    { "customer_id": "cust005", "user_id": "user001", "origin": "Sheffield, United Kingdom" }
  ]
}
```

#### Vessel JSON
```
{
  "id": "vessel001",
  "name": "Boaty McBoatFace",
  "max_weight": 200000,
  "capacity": 500
}
```

## Deployment

