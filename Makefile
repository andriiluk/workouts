logLvl = debug
run:
	export WORKOUTSVC_LOG_LEVEL=$(logLvl) && go run . run

lint:
	golangci-lint run

k8s-apply-muscle:
	kubectl apply -f ./k8s/musclesvc

argocd:
	kubectl port-forward -n argocd svc/argocd-server 9000:443