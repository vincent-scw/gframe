# Deploy to Azure Kubernetes Service (AKS)

## Steps
- Follow the [tutorials](https://docs.microsoft.com/en-us/azure/aks/) to setup AKS

### Kube-system
- Deploy [Kubernetes Dashboard](https://github.com/kubernetes/dashboard) to Kubernetes
  - Access http://localhost:8001/api/v1/namespaces/kube-system/services/kubernetes-dashboard/proxy/ instead. There is issue with the link in the page.
  - Apply ServiceAccount via command `kubectl apply -f https://raw.githubusercontent.com/vincent-scw/gframe/master/deploy/service-account.yaml`.
  - Use token revealed by `kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep admin-user | awk '{print $1}')` to login.
- Create Namespaces `kubectl apply -f https://raw.githubusercontent.com/vincent-scw/gframe/master/deploy/namespaces.yaml`

### Infrastructure
- Deploy Redis to Kubernetes
  - Run `kubectl apply -f https://raw.githubusercontent.com/vincent-scw/gframe/master/deploy/infra/redis.yaml` to deploy
  - There should be one master node and one backup (has not been added).
- Deploy Zookeeper to Kubernetes
  - Run `kubectl apply -f https://raw.githubusercontent.com/vincent-scw/gframe/master/deploy/infra/zookeeper.yaml` to deploy
- Deploy Kafka to Kubernetes
  - Run `kubectl apply -f https://raw.githubusercontent.com/vincent-scw/gframe/master/deploy/infra/kafka.yaml` to deploy
  - In the deployment, it contains 2 kafka brokers
  - After deployment, run `kubectl get pods -n infra`. Then find and copy a kafka pod like `kafka-1-deployment-{id}`
  - Get into the pod as `kubectl exec -it kafka-1-deployment-{id} -n infra /bin/bash`
  - Create topic by `kafka-topics.sh --create --zookeeper zookeeper-svc:2181 --replication-factor 1 --partitions 2 --topic player`
  - Confirm `Created topic player.`
- Deploy Kong to Kubernetes
  - Reference to [Kong Ingree on Azure Kubernetes Service](https://github.com/Kong/kubernetes-ingress-controller/blob/master/docs/deployment/aks.md)
  - Reference to [Kong Helm Chart](https://github.com/helm/charts/tree/master/stable/kong)
  - Run `helm install stable/kong --set ingressController.enabled=true --set proxy.loadBalancerIP=23.100.94.224 --set proxy.type=LoadBalancer --name kong --namespace infra`

### Applications
- Deploy gframe services
  - Run `kubectl apply -f https://raw.githubusercontent.com/vincent-scw/gframe/master/deploy/apps/` to deploy
