tidy:
	go mod tidy

run:
	go run cmd/main.go

build:
	go build cmd/main.go 

docker-build: 
	docker build -t simple-go-image:dev .

kind-create:
	kind create cluster \
		--name simple-cluster
		--config k8s/kind/kind-config.yaml

kind-destroy:
	kind delete cluster --name simple-cluster

kind-load:
	kind load docker-image simple-go-image:dev --name simple-cluster

# Generates a yaml using base and overlays
kustomize-build: 
	kustomize build k8s/simple-go-api/dev

# Pipe generated yaml to kubectl apply
kustomize-apply: 
	kustomize build k8s/simple-go-api/dev | kubectl apply -f -

skaffold-init:
	skaffold init

skaffold-run: 
	skaffold run

skaffold-dev: 
	skaffold dev
