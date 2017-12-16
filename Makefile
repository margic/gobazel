# creates the local development environment
.PHONY: create
create: config
	bazel run --cpu=k8 //deploy/local:create
	helm init

.PHONY: delete
delete:
	bazel run //deploy/local:delete


# Metrics - enable the metrics addon
.PHONY: applyMetrics
applyMetrics:
	helm install --namespace gobazel --name prometheus -f deploy/addons/prometheus-values.yaml stable/prometheus
	helm install --namespace gobazel --name grafana -f deploy/addons/grafana-values.yaml stable/grafana
	bazel run --cpu k8 //deploy/addons:apply-metrics
	bazel run --cpu k8 //dashboard:deploy-dashboard.apply


.PHONY: deleteMetrics
deleteMetrics:
	- bazel run --cpu k8 //deploy/addons:delete-metrics
	- bazel run --cpu k8 //dashboard:deploy-dashboard.delete
	- helm del --purge grafana
	- helm del --purge prometheus




# Tracing - enable the tracing addon
.PHONY: applyTracing
applyTracing:
	bazel run --cpu k8 //deploy/addons:apply-tracing.apply

.PHONY: deleteTracing
deleteTracing:
	bazel run --cpu k8 //deploy/addons:apply-tracing.delete


# Services deploys all the services
.PHONY: deploy
deploy:
	bazel run --cpu=k8 //deploy/devd

# syncs the system time with minikube
.PHONY: sync
sync:
	minikube ssh -- docker run -i --rm --privileged --pid=host debian nsenter -t 1 -m -u -n -i date -u $(date -u +%m%d%H%M%Y)

.PHONY: clean
clean:
	baezl clean

.PHONY: config
config:
	deploy/local/kubectl-config.sh
