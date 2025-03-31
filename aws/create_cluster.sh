eksctl create cluster -f aws/cluster.yaml
kubectl create secret generic db-credentials --from-env-file=secrets.env -n default
aws eks update-kubeconfig --name inventory-cluster --region us-east-1