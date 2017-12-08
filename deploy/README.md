Deploy contains the kubernetes config files

ingress-rbac.yml adds rolebase access control for ingress controller
ingress-deploy.yml deploys the ingress controller


# Output Prometheus Helm Install
```
NAME:   monitoring
LAST DEPLOYED: Thu Nov 23 08:28:54 2017
NAMESPACE: default
STATUS: DEPLOYED

RESOURCES:
==> v1/Pod(related)
NAME                                                      READY  STATUS             RESTARTS  AGE
monitoring-prometheus-node-exporter-9dmfc                 0/1    ContainerCreating  0         1s
monitoring-prometheus-alertmanager-5f6f465c6c-rfxbc       0/2    ContainerCreating  0         1s
monitoring-prometheus-kube-state-metrics-56d578f9b-rnpwh  0/1    ContainerCreating  0         1s
monitoring-prometheus-pushgateway-c896596b4-v2j4n         0/1    ContainerCreating  0         1s
monitoring-prometheus-server-7cb48db59d-wrqn5             0/2    Pending            0         1s

==> v1/ConfigMap
NAME                                DATA  AGE
monitoring-prometheus-alertmanager  1     1s
monitoring-prometheus-server        3     1s

==> v1/PersistentVolumeClaim
NAME                                STATUS  VOLUME                                    CAPACITY  ACCESS MODES  STORAGECLASS  AGE
monitoring-prometheus-alertmanager  Bound   pvc-046955d6-d063-11e7-8af5-7a9e95798150  2Gi       RWO           standard      1s
monitoring-prometheus-server        Bound   pvc-0469b863-d063-11e7-8af5-7a9e95798150  8Gi       RWO           standard      1s

==> v1/Service
NAME                                      TYPE       CLUSTER-IP  EXTERNAL-IP  PORT(S)   AGE
monitoring-prometheus-alertmanager        ClusterIP  10.0.0.125  <none>       80/TCP    1s
monitoring-prometheus-kube-state-metrics  ClusterIP  None        <none>       80/TCP    1s
monitoring-prometheus-node-exporter       ClusterIP  None        <none>       9100/TCP  1s
monitoring-prometheus-pushgateway         ClusterIP  10.0.0.66   <none>       9091/TCP  1s
monitoring-prometheus-server              ClusterIP  10.0.0.163  <none>       80/TCP    1s

==> v1beta1/DaemonSet
NAME                                 DESIRED  CURRENT  READY  UP-TO-DATE  AVAILABLE  NODE SELECTOR  AGE
monitoring-prometheus-node-exporter  1        1        0      1           0          <none>         1s

==> v1beta1/Deployment
NAME                                      DESIRED  CURRENT  UP-TO-DATE  AVAILABLE  AGE
monitoring-prometheus-alertmanager        1        1        1           0          1s
monitoring-prometheus-kube-state-metrics  1        1        1           0          1s
monitoring-prometheus-pushgateway         1        1        1           0          1s
monitoring-prometheus-server              1        1        1           0          1s


NOTES:
The Prometheus server can be accessed via port 80 on the following DNS name from within your cluster:
monitoring-prometheus-server.default.svc.cluster.local


Get the Prometheus server URL by running these commands in the same shell:
  export POD_NAME=$(kubectl get pods --namespace default -l "app=prometheus,component=server" -o jsonpath="{.items[0].metadata.name}")
  kubectl --namespace default port-forward $POD_NAME 9090


The Prometheus alertmanager can be accessed via port 80 on the following DNS name from within your cluster:
monitoring-prometheus-alertmanager.default.svc.cluster.local


Get the Alertmanager URL by running these commands in the same shell:
  export POD_NAME=$(kubectl get pods --namespace default -l "app=prometheus,component=alertmanager" -o jsonpath="{.items[0].metadata.name}")
  kubectl --namespace default port-forward $POD_NAME 9093


The Prometheus PushGateway can be accessed via port 9091 on the following DNS name from within your cluster:
monitoring-prometheus-pushgateway.default.svc.cluster.local


Get the PushGateway URL by running these commands in the same shell:
  export POD_NAME=$(kubectl get pods --namespace default -l "app=prometheus,component=pushgateway" -o jsonpath="{.items[0].metadata.name}")
  kubectl --namespace default port-forward $POD_NAME 9093

For more information on running Prometheus, visit:
https://prometheus.io/
```

# Grafana helm output 
```
➜  gobazel git:(master) ✗ helm --namespace gobazel --name grafana install stable/grafana
NAME:   grafana
LAST DEPLOYED: Thu Nov 23 09:09:55 2017
NAMESPACE: gobazel
STATUS: DEPLOYED

RESOURCES:
==> v1/Secret
NAME             TYPE    DATA  AGE
grafana-grafana  Opaque  2     1s

==> v1/ConfigMap
NAME                    DATA  AGE
grafana-grafana-config  1     1s
grafana-grafana-dashs   0     1s

==> v1/PersistentVolumeClaim
NAME             STATUS  VOLUME                                    CAPACITY  ACCESS MODES  STORAGECLASS  AGE
grafana-grafana  Bound   pvc-bf091c5d-d068-11e7-8af5-7a9e95798150  1Gi       RWO           standard      1s

==> v1/Service
NAME             TYPE       CLUSTER-IP  EXTERNAL-IP  PORT(S)  AGE
grafana-grafana  ClusterIP  10.0.0.120  <none>       80/TCP   1s

==> v1beta1/Deployment
NAME             DESIRED  CURRENT  UP-TO-DATE  AVAILABLE  AGE
grafana-grafana  1        1        1           0          1s

==> v1/Pod(related)
NAME                              READY  STATUS   RESTARTS  AGE
grafana-grafana-6f76cdd8df-wll9m  0/1    Pending  0         1s


NOTES:
1. Get your 'admin' user password by running:

   kubectl get secret --namespace gobazel grafana-grafana -o jsonpath="{.data.grafana-admin-password}" | base64 --decode ; echo

2. The Grafana server can be accessed via port 80 on the following DNS name from within your cluster:

   grafana-grafana.gobazel.svc.cluster.local

   Get the Grafana URL to visit by running these commands in the same shell:

     export POD_NAME=$(kubectl get pods --namespace gobazel -l "app=grafana-grafana,component=grafana" -o jsonpath="{.items[0].metadata.name}")
     kubectl --namespace gobazel port-forward $POD_NAME 3000

3. Login with the password from step 1 and the username: admin

```

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