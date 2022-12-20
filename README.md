# minikube on Linux arm64

minikube setup based on ubuntu-20.04-arm64 running as a Vagrant Virtual Machine with the _Parallels_ provider.

## Resources
 - Memory: 2.75 GB
 - CPU: 2

## Applications
- Docker
- Docker Compose
- minikube

## Just start

### Start the Virtual Machine
```
vagrant up
```
### SSH into the Virtual Machine 
```
vagrant ssh
```

### Launch minikube
```
minikube start --memory=2gb --kubernetes-version=v1.24.7 
```

### Install kubectl
```
curl -LO https://dl.k8s.io/release/v1.24.7/bin/linux/arm64/kubectl
```

```
curl -LO https://dl.k8s.io/release/v1.24.7/bin/linux/arm64/kubectl.sha256
echo "$(cat kubectl.sha256)  kubectl" | sha256sum --check
```

```
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
```

## Operations

### Get all running pods!
```
kubectl get pods -A
```

### Expose Kubernetes API using kubectl proxy
```
kubectl proxy --port=8080 --address=0.0.0.0 --accept-hosts='.+'
```

#### Test Kubernetes API from the host machine
```
curl http://localhost:8080/apis
```
