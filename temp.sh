  
kubectl create secret docker-registry gcr-json-key \
--docker-server=https://gcr.io \
--docker-username=_json_key \
--docker-password="$(cat ~/gcr-json-key.json)" \
--docker-email=k8gcrimagepuller@my-test-project-167717.iam.gserviceaccount.com
