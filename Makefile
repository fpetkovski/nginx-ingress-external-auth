.ONESHELL:

TAG=0.1
PREFIX?=electroma/ingress-demo-
ARCH?=amd64
GOLANG_VERSION=1.9
TEMP_DIR:=$(shell mktemp -d)

.PHONY: build
build: 
	eval $$(minikube docker-env)
	docker build -t auth-service:$(TAG) auth-service 			  -f auth-service/Dockerfile
	docker build -t echo-service:$(TAG) echo-service 			  -f echo-service/Dockerfile
	docker build -t auth-proxy:$(TAG)   auth-proxy/proxy 		  -f auth-proxy/proxy/Dockerfile
	docker build -t auth-init:$(TAG)    auth-proxy/init-container -f auth-proxy/init-container/Dockerfile

clean:
	kubectl delete -f deploy/ --wait=true
	eval $$(minikube docker-env)
	docker rmi \
		auth-service:$(TAG) \
		echo-service:$(TAG) \
		auth-proxy:$(TAG) \
		auth-init:$(TAG)

.PHONY: deploy
deploy: build
	kubectl apply -f deploy/

redeploy: clean deploy