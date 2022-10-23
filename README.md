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

### Get running pods!
```
minikube kubectl -- get pods -A
```