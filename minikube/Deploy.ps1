
$ProfileName = "demo-cluster"
$cpuCount = 4
$memory = 30*1024

function StartCluster
{
    & minikube start -p $ProfileName --addons ingress --cpus=$cpuCount --memory $memory --docker-opt=log-driver=json-file --docker-opt=log-opt=max-size=10M --docker-opt=log-opt=max-file=5
}

# start the different clusters minikube with profile
$UserName = [System.Environment]::UserName
StartCluster -clusterName $UserName