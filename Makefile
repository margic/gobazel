.PHONY: clean
clean:

.PHONY: build
build:

.PHONY: helm
helm:
	#helm init
	helm install --namespace gobazel --name prometheus stable/prometheus
	helm install --namespace gobazel --name grafana stable/grafana

.PHONY: config
config:
	scripts/kubectl-config.sh