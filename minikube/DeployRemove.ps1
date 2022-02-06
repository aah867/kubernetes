$ProfileName = "demo-cluster"

& minikube stop -p $ProfileName
& minikube delete -p $ProfileName

# remove all old docker images
& docker system prune --all --force