# Infrastructure

Project uses minikube as 'production' deployment

## Minikube

### Docker images

You will need to follow these steps for the service images to be available in Minikube
```
# Start minikube
minikube start

# Set docker env
eval $(minikube docker-env)

# Build image
docker build -t foo:0.0.1 .

# Run in minikube
kubectl run hello-foo --image=foo:0.0.1 --image-pull-policy=Never

# Check that it's running
kubectl get pods
```

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


