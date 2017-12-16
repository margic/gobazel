# Deploy
Deploy contains the deployment files for the local environment and deployment environments.

## Folder Structure

* addons - contains optional features that can be added to the local developer environment
* dev - kubernetes deployment files for development runtime environment
* qa - kubernetes deployment files for qa environment
* local - kubernetes deployment files for standing up local developer environment in minikube
* templates - template files used to generate dev/qa/prod configurations. TODO see issue #6

## Addons
Addons are provided to make developing necessary features such as dashboards and tracing eaier.
Services should not depend on the addons for running. They are intended to be enabled to support
development of features that support services.

Running all the addons would cripple the average local machine. Addons are intetend to be used
to support development activity and can be enabled and disabled at will.

The following addons are provided
* Metrics - Prometheus an Grafana configured to discover services.
* Tracing - Jaeger tracing to provide view over call paths.
* Logging - TODO see issue #7

### Metrics
Metrics addon provides prometheus and grafana for local development of instrumentation and 
dashboards

`make applyMetrics`
`make deleteMetrics`

### Tracing
Tracing addon provides all-in-one jaeger tracing. This uses a memory store for jaeger and it's use
should be limited to local testing.

`make applyTracing`
`make deleteTracing`

### Logging
TODO


NOTES:
Grafana user admin, password defaults to Passw0rd


registry.yaml depends on ingress being setup
setups registry service and ingress

to allow access to the registry make sure to forward port 5000 to localhost.
Bazel can currently only push to insecure registry on localhost.

By using localhost:5000 as the registry address allows local docker daemon (optional)
and the minikube docker daemon to access the registry deployed in the namespace

```
kubectl port-forward --namespace gobazel \
$(kubectl get po -n gobazel | grep registry-proxy | \
awk '{print $1;}') 5000:80
```

kubectl get pods -a -n gobazel -o json  | jq -r '.items[] | select(.status.phase == "Running") or ([ .status.conditions[] | select(.type == "Ready" and .status == "False") ] | length ) == 1 ) | .metadata.namespace + "/" + .metadata.name'