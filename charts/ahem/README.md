# ahem

A Helm chart for Ahem

![Version: 0.0.3](https://img.shields.io/badge/Version-0.0.3-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.0.3](https://img.shields.io/badge/AppVersion-0.0.3-informational?style=flat-square)

## Installing the Chart

```bash
helm upgrade --install --namespace default --create-namespace \
  ahem oci://public.ecr.aws/brandonwagner/ahem \
  --version 0.0.3
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| additionalAnnotations | object | `{}` | Additional annotations to add into metadata. |
| additionalLabels | object | `{}` | Additional labels to add into metadata. |
| affinity | object | `{}` | Affinity rules for scheduling the pod. |
| env | list | `[]` | Additional environment variables for the controller pod. |
| fullnameOverride | string | `""` | Overrides the chart's computed fullname. |
| hostNetwork | bool | `false` | Bind the pod to the host network. |
| image | string | `"public.ecr.aws/brandonwagner/ahem:v0.0.3@sha256:83652195abe21dca110b20f20a20161aaa34c9ff117489e08b503b5586e732e0"` | Controller image. |
| imagePullPolicy | string | `"IfNotPresent"` | Image pull policy for Docker images. |
| imagePullSecrets | list | `[]` | Image pull secrets for Docker images. |
| nameOverride | string | `""` | Overrides the chart's name. |
| nodeSelector | object | `{"kubernetes.io/os":"linux"}` | Node selectors to schedule the pod to nodes with labels. |
| podAnnotations | object | `{}` | Additional annotations for the pod. |
| podLabels | object | `{}` | Additional labels for the pod. |
| podSecurityContext | object | `{"fsGroup":1000}` | SecurityContext for the pod. |
| priorityClassName | string | `nil` | PriorityClass name for the pod. |
| replicas | int | `1` | Number of replicas. |
| resources | object | `{"limits":{"memory":"128Mi"},"requests":{"cpu":"100m","memory":"128Mi"}}` | Resources for the pods. |
| securityContext | object | `{}` | SecurityContext for the controller container. |
| strategy | object | `{"type":"Recreate"}` | Strategy for updating the pod. |
| terminationGracePeriodSeconds | string | `nil` | Override the default termination grace period for the pod. |
| tolerations | list | `[]` | Tolerations to allow the pod to be scheduled to nodes with taints. |
| topologySpreadConstraints | list | `[]` | topologySpreadConstraints to increase the controller resilience |

