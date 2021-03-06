# Infrastructure

Project uses minikube as 'production' deployment

## Minikube

### Addons
#### Ingress

From the following [resource](https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/)

You must install the following Minikube addons
```
minikube addons enable ingress
```

Create the ingress service
```
kubectl apply -f ingress.yml
```

Update `/etc/hosts` with the assigned ip and use the host name specified in the ingress definition file e.g:
```
172.1.2.3 ingress.local
```

You can now visit one URL and reach various services via the specfied paths!
_Web_ : `http://ingress.local/`
_API_ : `http://ingrses.local/api`

### EXTERNAL IPs
***This might not be necessary now we have the ingress controller***

This is important, as we are running Minikube we will never get external ips assigned to our services that request them, to work around this we must run the following against a running Minikube cluster
#### Network setup
```
sudo ip route add $(cat ~/.minikube/profiles/minikube/config.json | jq -r ".KubernetesConfig.ServiceCIDR") via $(minikube ip)
kubectl run minikube-lb-patch --replicas=1 --image=elsonrodriguez/minikube-lb-patch:0.1 --namespace=kube-system
```
#### Network teardown
```
kubectl delete deployment minikube-lb-patch -nkube-system
sudo ip route delete $(cat ~/.minikube/profiles/minikube/config.json | jq -r ".KubernetesConfig.ServiceCIDR") via $(minikube ip)
```

See the following for more information:
1. https://github.com/knative/serving/blob/b31d96e03bfa1752031d0bc4ae2a3a00744d6cd5/docs/creating-a-kubernetes-cluster.md#loadbalancer-support-in-minikube
2. https://github.com/elsonrodriguez/minikube-lb-patch

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

### Manually create a secret for db passwords.. e.g:
```
kubectl create secret generic postgres --from-literal=password=postgres
```

### Database
```
kubectl create -f ./deployments/mongodb-ssd.yml
kubectl create -f ./deployments/mongodb-deployment.yml
kubectl create -f ./deployments/mongodb-service.yml
```

### Services
Each service has its own deployment script in its own deployment folder


