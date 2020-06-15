# Vessel Service

This is the Vessel service

Generated with

```
micro new vessel --namespace=go.micro --type=service
```

## Usage

A Makefile is included for convenience

### Build the binary

```
make build
```

### Run the service
```
./vessel-service
```

### Build a docker image
If you are building the image for use in Minikube then first run
```
eval (minikube docker-env)
```

```
make docker
```

### Deploy to K8s Minikube
```
make deploy
```

### Command line call (via micro api)
```
curl ingress.local/api/rpc -XPOST -d '{
  "request": { 
    "name": "test", 
    "capacity": 200, 
    "max_weight": 100000, 
    "available": true 
  },
  "method": "VesselService.Create",
  "service": "com.foo.service.vessel"
}' -H 'Content-Type: application/json'
```
