# Infrastructure

Project uses minikube as 'production' deployment

## Minikube

### Start

```
minikube start --vm-driver=kvm2 --memory=4096
```

## Deploy

### Database
```
kubectl create -f ./deployments/mongodb-ssd.yml
kubectl create -f ./deployments/mongodb-deployment.yml
kubectl create -f ./deployments/mongodb-service.yml
```

### Services
Each service has its own deployment script


