# Deploy to Azure Kubernetes Service (AKS)

## Steps
- Follow the [tutorials](https://docs.microsoft.com/en-us/azure/aks/) to setup AKS
- Install Helm and Tiller
  - `helm init --service-account tiller --history-max 200`

### Kube-system
- [Depreciated, Kubernetes Dashboard is included in AKS by default] Deploy [Kubernetes Dashboard](https://github.com/kubernetes/dashboard) to Kubernetes
  - Reference to https://docs.microsoft.com/en-us/azure/aks/kubernetes-dashboard
- Create Namespaces `kubectl apply -f ./namespaces.yaml`
- Add ServiceAccount via `kubectl apply -f ./service-account.yaml`
  - `tiller` is for Helm
  - `kubernetes-dashboard` is for Kubernetes Dashboard

### Infrastructure
- Deploy [Redis](https://redis.io/) to Kubernetes
  - Run `kubectl apply -f ./infra/redis.yaml` to deploy
  - There should be one master node and one backup (has not been added).
- Deploy [Zookeeper](https://zookeeper.apache.org/) to Kubernetes
  - Run `kubectl apply -f ./infra/zookeeper.yaml` to deploy
- Deploy [Kafka](https://kafka.apache.org/) to Kubernetes
  - Run `kubectl apply -f ./infra/kafka.yaml` to deploy
  - In the deployment, it contains 2 kafka brokers
  - After deployment, run `kubectl get pods -n infra`. Then find and copy a kafka pod like `kafka-1-deployment-{id}`
  - Get into the pod as `kubectl exec -it kafka-1-deployment-{id} -n infra /bin/bash`
  - Create topic by `kafka-topics.sh --create --zookeeper zookeeper-svc:2181 --replication-factor 1 --partitions 2 --topic player`
  - Confirm `Created topic player.`
- Deploy [Nginx](https://www.nginx.com/) to Kubernetes
  - Create a static IP, reference to https://docs.microsoft.com/en-us/azure/aks/static-ip
  - Run `helm install stable/nginx-ingress --name nginx-gframe --namespace gframe --set controller.replicaCount=2 --set controller.service.loadBalancerIP={{StaticIP}}`
 
### Monitoring
- Deploy [Promethuse](https://prometheus.io/) to Kubernetes
  - Run `helm install --name prometheus --namespace monitoring stable/prometheus-operator -f ./monitoring/prometheus-settings.yaml` to deploy
  - After Prometheus installed, run `kubectl apply -f ./monitoring/services.yaml` to add services to configuration
- Deploy [Jaeger](https://www.jaegertracing.io/) to Kubernetes

### Applications
- Deploy gframe services
  - Run `kubectl apply -f ./apps/` to deploy
