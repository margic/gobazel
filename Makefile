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
	bazel run --cpu k8 //deploy/addons:deploy-metrics.apply
	bazel run --cpu k8 //dashboard:deploy-dashboard.apply

# NATS
.PHONY: applyNats
applyNats:
	bazel run --cpu k8 //deploy/addons:deployNats
		
.PHONY: deleteNats
deleteNats:
	bazel run --cpu k8 //deploy/addons:deleteNats
	
.PHONY: deleteMetrics
deleteMetrics:
	- bazel run --cpu k8 //deploy/addons:deploy-metrics.delete
	- bazel run --cpu k8 //dashboard:deploy-dashboard.delete
	- helm del --purge grafana
	- helm del --purge prometheus

# Tracing - enable the tracing addon
.PHONY: applyTracing
applyTracing:
	bazel run --cpu k8 //deploy/addons:deploy-tracing.apply

.PHONY: deleteTracing
deleteTracing:
	bazel run --cpu k8 //deploy/addons:deploy-tracing.delete

# Logging - enable the logging addon
.PHONY: applyLogging
applyLogging:
	bazel run --cpu k8 //deploy/addons:deployLogging

.PHONY: deleteLogging
deleteLogging:
	bazel run --cpu k8 //deploy/addons:deleteLogging

# Services deploys all the services
.PHONY: applyServices
applyServices:
	- bazel run --cpu=k8 //deploy/dev:deploy-greeting.apply
	- bazel run --cpu=k8 //deploy/dev:deploy-greet.apply
	- bazel run --cpu=k8 //deploy/dev:deploy-launcher.apply

.PHONY: deleteServices
deleteServices:
	- bazel run --cpu=k8 //deploy/dev:deploy-greeting.delete
	- bazel run --cpu=k8 //deploy/dev:deploy-greet.delete
	- bazel run --cpu=k8 //deploy/dev:deploy-launcher.delete



#############################
# Some useful utility targets
#############################


# syncs the system time with minikube
.PHONY: sync
sync:
	minikube ssh -- docker run -i --rm --privileged --pid=host debian nsenter -t 1 -m -u -n -i date -s @$(shell date -u +%s)

#$(date -u +'%s')

.PHONY: clean
clean:
	bazel clean
	rm -rf bazel-*

# Configures the workspace kube folder with certs from minikube
.PHONY: config
config:
	deploy/local/kubectl-config.sh

# Genereate protobuf files manually
.PHONY: generate
generate:
	protoc -I protos/ --go_out=plugins=grpc:protos protos/*.proto
	bazel run //:gazelle