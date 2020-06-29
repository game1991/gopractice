.PHONY : test

CURDIR=$(shell pwd)
DOCKER_TARGET= hub.deepin.com/goweb-edu/ganlei
DOCKER_BUILDER_TARGET=$(DOCKER_TARGET).builder
TARGET=projectdemo

build:
	go build -mod vendor -v -o $(CURDIR)/$(TARGET) cmd/main.go

run: build
	SERVER_ROOT=$(CURDIR) $(CURDIR)/$(TARGET)

docker-build:
	docker image inspect golang >/dev/null || docker pull golang
	docker image inspect debian >/dev/null || docker pull debian
	DOCKER_BUILDKIT=1 docker build -f deployment/Dockerfile --target builder -t $(DOCKER_BUILDER_TARGET) .

docker:
	DOCKER_BUILDKIT=1 docker build -f deployment/Dockerfile -t $(DOCKER_TARGET) .

docker-release:
	docker push $(DOCKER_BUILDER_TARGET)
	docker push $(DOCKER_TARGET)

docker-push:
	docker push $(DOCKER_BUILDER_TARGET)
	docker push $(DOCKER_TARGET)

migrate:
	echo SERVER_ROOT=$(CURDIR)

test:
	echo cd $(CURDIR)/test; SERVER_ROOT=$(CURDIR) ginkgo -v -r -cover

clean:
	rm $(CURDIR)/$(TARGET)