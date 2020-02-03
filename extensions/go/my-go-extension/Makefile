RELEASE?=0.0.1
APP?=my-go-extension
GOOS?=linux
GOARCH?=amd64
CONTAINER_IMAGE?=gabbi/${APP}:${RELEASE}
NAMESPACE?=dev

clean:
	rm -rf ${APP}

vet:
	GO111MODULE=on go vet .

test:
	GO111MODULE=on go test -v -race ./...

build: clean vet test
	GO111MODULE=on CGO_ENABLED=0 GOARCH=${GOARCH} go build \
		-ldflags "-s -w" \
		-o ${APP}

build-image: build
	docker build -t $(CONTAINER_IMAGE) .

push-image:
	docker push $(CONTAINER_IMAGE)

tidy:
	GO111MODULE=on go mod tidy

create-service:
	kubectl -n $(NAMESPACE) create deployment $(APP) --image=$(CONTAINER_IMAGE) --dry-run -o yaml > deployment/deployment-ohne-port.yaml
	kubectl patch -f ./deployment/deployment-ohne-port.yaml -p '{"spec": {"template" :{"spec":{"containers":[{"name": "$(APP)", "ports":[{"containerPort": 8080}]}]}}}}' --dry-run  -o yaml > deployment/deployment.yaml
	kubectl expose -f deployment/deployment.yaml --type=ClusterIP --dry-run -o yaml > deployment/service.yaml
	rm deployment/deployment-ohne-port.yaml

expose-service:
	../../../scripts/create-api.sh "./deployment/api.yaml" $(APP) $(NAMESPACE)
