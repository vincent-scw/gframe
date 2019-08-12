# Deploy to Azure Kubernetes Service (AKS)

## Steps
- Follow the [tutorials](https://docs.microsoft.com/en-us/azure/aks/) to setup AKS
- Deploy [Kubernetes Dashboard](https://github.com/kubernetes/dashboard) to Kubernetes
  - Access http://localhost:8001/api/v1/namespaces/kube-system/services/kubernetes-dashboard/proxy/ instead. There is issue with the link in the page.
  - Apply ServiceAccount via command `kubectl apply -f https://raw.githubusercontent.com/vincent-scw/gframe/master/deploy/service-account.yaml`.
  - Use token revealed by `kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep admin-user | awk '{print $1}')` to login.
