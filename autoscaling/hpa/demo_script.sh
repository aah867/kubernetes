#!/bin/bash

########################################## HPA ##################################

# Step-0: Prerequisites: Kubernetes cluster, Kubectl, Metrics server
minikube profile list
minikube -p demo-cluster addons list
minikube -p demo-cluster addons enable metrics-server
minikube -p demo-cluster addons disable metrics-server
kubectl get pod,svc -n kube-system

# Step 1: Run and expose php-apache server

# Step 2: Deploy Horizontal Pod Autoscaler

# Step 3: Monitor HPA and CPU utilization of Demo-app
kubectl get hpa --watch
kubectl get deployment --watch
watch kubectl top pods

# Step 4: Increase load
kubectl run -i --tty load-generator --image=busybox /bin/sh
while true; do wget -q -O- http://php-apache.default.svc.cluster.local; done

# Step 5: Cleanup:
kubectl delete po load-generator
#minikube -p demo-cluster addons disable metrics-server

########################################## VPA ##################################

# Step 2: Download and Deploy the Vertical Pod Autoscaler
/home/kratos/aah_github/kubernetes/autoscaling/vpa/installer/hack/vpa-up.sh
kubectl get po -n kube-system

# Step 3: Deploy the Sample Application
kubectl apply -f demo-app.yaml

# Step 4: Deploy vpa to sample app
kubectl apply -f /home/kratos/aah_github/kubernetes/autoscaling/vpa/vpa.yaml

# Step 5: Monitor app and vpa:
kubectl describe vpa
watch 'kubectl describe po demo-app-vpa'
kubectl get event -n kube-system --watch
