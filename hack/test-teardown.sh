set +e

kubectl delete -f ./examples/database
kubectl delete -f ./examples/vault
kubectl delete -f ./examples/providerconfig
kubectl delete -f ./package/crds
