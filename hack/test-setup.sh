set +e

kubectl apply -f ./package/crds
kubectl apply -f ./examples/providerconfig
kubectl apply -f ./examples/vault
kubectl apply -f ./examples/database
