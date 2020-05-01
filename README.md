# REMORA
Dockerhub-webhook trigger

## Curl
```
curl --request GET \
  --url 'http://localhost/v1/playbook?apps=example&inventory=127.0.0.1&connection=local&report=0'
```
report = 1 for syncronous mode