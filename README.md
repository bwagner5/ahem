# Ahem

Ahem is a simple application that waits for a SIGTERM signal and then waits a configurable period of time before exiting after receiving the SIGTERM. This application is used to demonstrate how Kubernetes pods shutdown in response to an eviction.

## Installation:


```bash
helm upgrade --install --namespace default --create-namespace ahem oci://public.ecr.aws/brandonwagner/ahem --version 0.0.2
```

Only install on Spot Instances and increase the delay to breach the default termination grace period:

```bash
helm upgrade --install --namespace default --create-namespace ahem oci://public.ecr.aws/brandonwagner/ahem --version 0.0.2 \
    --set nodeSelector."karpenter\.sh/capacity-type=spot" \
    --set 'env[0].name=delay' \
    --set 'env[0].value=60s'
```