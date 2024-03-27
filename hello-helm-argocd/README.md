
helm repo update 	# Make sure we get the latest list of charts
helm show values bitnami/mysql # show values of chart to find params to override
helm install bitnami/mysql --generate-name
helm install --set primary.service.type=NodePort --set primary.service.nodePorts.mysql=6603 bitnami/mysql --generate-name --dry-run --debug
helm get manifest mysql-1711557852 # show manifest of running pod to find username/password
kubectl get secrets/mysql-1711557852 -o json # show username/password of mysql in base64 encoded
echo <encoded base64 password> | base64 --decode # decode base64
