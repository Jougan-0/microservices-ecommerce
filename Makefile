.PHONY: all docker k8s helm clean

all: docker

docker:
	docker-compose up --build

k8s:
	kubectl apply -f k8s/

helm:
	helm install ecommerce ./helm

clean:
	docker-compose down
	kubectl delete namespace ecommerce
	helm uninstall ecommerce
