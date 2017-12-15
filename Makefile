.PHONY: clean
clean:

.PHONY: build
build:

.PHONY: helm
helm:
	- helm install --namespace gobazel --name prometheus -f deploy/prometheus-values.yaml stable/prometheus
	- helm install --namespace gobazel --name grafana -f deploy/grafana-values.yaml stable/grafana
	- kubectl get secret --namespace gobazel grafana-grafana -o jsonpath="{.data.grafana-admin-password}" | base64 --decode ; echo

.PHONY: config
config:
	scripts/kubectl-config.sh

# creates the local development environment
.PHONY: create
create:
	bazel run --cpu=k8 //deploy:create

# deletes the local development namespace
.PHONY: delete
delete:
	- helm del --purge grafana
	- helm del --purge prometheus
	bazel run //deploy:delete

# deploy services
.PHONY: deploy
deploy:
	bazel run --cpu=k8 //deploy:deploy

# syncs the system time with minikube
.PHONY: sync
sync:
	minikube ssh -- docker run -i --rm --privileged --pid=host debian nsenter -t 1 -m -u -n -i date -u $(date -u +%m%d%H%M%Y)
