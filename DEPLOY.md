# Deploy to Azure Kubernetes Service (AKS)

## Steps
- Follow the [tutorials](https://docs.microsoft.com/en-us/azure/aks/) to setup AKS
- Deploy [Kubernetes Dashboard](https://github.com/kubernetes/dashboard) to Kubernetes
  - Access http://localhost:8001/api/v1/namespaces/kube-system/services/kubernetes-dashboard/proxy/ instead. There is issue with the link in the page.
