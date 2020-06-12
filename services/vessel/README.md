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
