kubectl apply -f db-claim0-persistentvolumeclaim.yaml
kubectl apply -f db-deployment.yaml
kubectl apply -f db-service.yaml
kubectl apply -f app-claim0-persistentvolumeclaim.yaml
kubectl apply -f app-deployment.yaml
kubectl apply -f app-service.yaml