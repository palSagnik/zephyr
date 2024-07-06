#!/bin/bash

set +x

cd ./kubernetes/init
kubectl apply -f db-volume.yml

# initializing secrets
cd ../configuration
kubectl apply -f app-secrets.yml

# initializing applications
cd ../definition
kubectl apply -f db-main.yml