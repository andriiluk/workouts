run:
	go run . run

lint:
	golangci-lint run

k8s-apply-muscle:
	kubectl apply -f ./k8s/musclesvc