# minikube on Linux arm64

minikube setup based on ubuntu-20.04-arm64 running as a Vagrant Virtual Machine with the _Parallels_ provider.

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
minikube start --kubernetes-version=v1.24.7 
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

### Create a job
```
kubectl create job hello --image nixery.dev/arm64/shell -- echo "Hello World"
```

### Expose Kubernetes API using kubectl proxy
```
kubectl proxy --port=8080 --address=0.0.0.0 --accept-hosts='.+'
```

#### Test Kubernetes API from the host machine
```
curl http://localhost:8080/apis
```

## Programming (Go)
The [src](https://github.com/jrhrmsll/minikube-arm64/tree/main/src) directory contains an simple program using the [client-go](https://github.com/kubernetes/client-go) library, to talk with the Kubernetes API
and list the existing jobs in the default namespace.

To create the program binary run (from the `src` directory):
```
GOOS=linux go build -o jobs main.go
```

in the **host** machine and execute the following in the **guess** one:

```
/vagrant/src/jobs
```

the output should be similar to:
```
hello
```

the name of the Job created before in the `default` namespace.
