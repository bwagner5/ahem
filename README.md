# Ahem

Ahem is a simple application that waits for a SIGTERM signal and then waits a configurable period of time before exiting after receiving the SIGTERM. This application is used to demonstrate how Kubernetes pods shutdown in response to an eviction.

## Installation:

```bash
cat <<EOF | kubectl apply -f -
 apiVersion: apps/v1
 kind: Deployment
 metadata:
   name: ahem
   labels:
     app: ahem
 spec:
   selector:
     matchLabels:
       app: ahem
   replicas: 0
   template:
     metadata:
       labels:
         app: ahem
     spec:
       containers:
       - image: public.ecr.aws/brandonwagner/ahem:v0.0.3
         name: ahem
         resources:
           requests:
             cpu: "1"
             memory: 256M
       topologySpreadConstraints:
       - labelSelector:
           matchLabels:
             app: ahem
         maxSkew: 1
         topologyKey: topology.kubernetes.io/zone
         whenUnsatisfiable: DoNotSchedule
EOF
```


```bash
helm upgrade --install --namespace default --create-namespace ahem oci://public.ecr.aws/brandonwagner/ahem --version 0.0.3
```

Only install on Spot Instances and increase the delay to breach the default termination grace period:

```bash
helm upgrade --install --namespace default --create-namespace ahem oci://public.ecr.aws/brandonwagner/ahem --version 0.0.3 \
    --set nodeSelector."karpenter\.sh/capacity-type=spot" \
    --set 'env[0].name=delay' \
    --set 'env[0].value=60s'
```
