RELEASE?=0.0.1
APP?=my-java-extension
DOCKER_ACCOUNT?=gabbi
CONTAINER_IMAGE?=${DOCKER_ACCOUNT}/${APP}:${RELEASE}
NAMESPACE?=dev

.PHONY: build build-docker clean tag push

all: clean build build-docker tag

clean:
	./gradlew clean
build:
	./gradlew -Pversion=$(RELEASE) -Pname=$(APP) clean build

build-image: clean build
	cp build/libs/$(APP)-$(RELEASE).jar docker
	docker build -t $(CONTAINER_IMAGE) --no-cache --rm docker
	rm docker/$(APP)-*.jar

push-image:
	docker push $(CONTAINER_IMAGE)

create-service:
	kubectl -n $(NAMESPACE) create deployment $(APP) --image=$(CONTAINER_IMAGE) --dry-run -o yaml > deployment/deployment-ohne-port.yaml
	kubectl patch -f ./deployment/deployment-ohne-port.yaml -p '{"spec": {"template" :{"spec":{"containers":[{"name": "$(APP)", "ports":[{"containerPort": 8080}]}]}}}}' --dry-run  -o yaml > deployment/deployment.yaml
	kubectl expose -f deployment/deployment.yaml --type=ClusterIP --dry-run -o yaml > deployment/service.yaml
	rm deployment/deployment-ohne-port.yaml

expose-service:
	../../../scripts/create-api.sh "./deployment/api.yaml" $(APP) $(NAMESPACE)