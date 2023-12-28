# Simple Deployment Services with Kubernetes

This is demo of deploying a REST application and MySQL Database in Kubernetes with <b>Minikube</b>

<br>

## Requirements
Install Docker, Kubernetes and Minikube by following this guide:  
https://programmingpercy.tech/blog/learn-kubernetes-the-easy-way/

<br>

## Commands
Assuming you already install above, test these following commands..  
<br>
### set docker context with Minikube context
```bash
docker context use default
```
### start Minikube  
```bash
minikube start
```  
### use Minikube docker environment  
This command need to be re-execute after each terminal restart.<br>your list of docker images and containe will be different after use this command..
<br>
<br>
linux or git bash:
```bash
eval $(minikube docker-env)
```  
windows:
```bash
minikube -p minikube docker-env | Invoke-Expression
```
<br>

### build image
```bash
docker build -t programmingpercy/hellogopher:1.0 .
```

### create resource (deployment, namespace, etc)
```bash
kubectl create -f <filepath or directory>
```
or
```bash
kubctl apply -f <filepath or directory>
```
<br>

### expose deployment
```bash
kubectl expose deployment <resourcename> --type=NodePort --port=8080
```

### list all resource
```bash
kubectl get all
```

### list/delete deployments/pod/service
```bash
kubectl <get/delete> <deployments/pods/service>
```

### visit the service (exposed deployment), this will open browser
```bash
minikube service <resourcename>
```

### enter container
```bash
kubectl exec <podname> -it -- bash
```

### set namespace
```bash
kubectl config set-context --current --namespace=<namespaceName>
```

### visit service with namespace
```bash
minikube service hellogopher -n hellogopher
```