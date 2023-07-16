PKG_NAME = ezone.xxxxx.com/xxxxx/xxxxx/alarm
COMMIT := $(shell git rev-parse --short HEAD)
VERSION := $(shell git describe --tags)
BuildTime := $(shell git show -s --format=%cd)
SERVER_DIR := "."

# IMAGE_VERSION := $(COMMIT)

build-image: #proto vendor
	DOCKER_BUILDKIT=0 docker build \
	--build-arg COMMIT=$(COMMIT) \
	--build-arg VERSION=$(VERSION) \
	--build-arg SERVER_DIR="$(SERVER_DIR)" \
	--build-arg PKG_NAME="$(PKG_NAME)" \
	--target base -t harbor.inner.galaxy.xxxxx.com/xxxxx/alarmv2:$(VERSION) ./
	docker push harbor.inner.galaxy.xxxxx.com/xxxxx/alarmv2:$(VERSION)
	docker tag harbor.inner.galaxy.xxxxx.com/xxxxx/alarmv2:$(VERSION) harbor.inner.galaxy.xxxxx.com/xxxxx/alarmv2:latest
	docker push harbor.inner.galaxy.xxxxx.com/xxxxx/alarmv2:latest

	for i in $(shell docker ps -a | grep Exited | awk '{print $$1}'); do if [ -n $$i ]; then docker rm $$i;fi done
	for i in $(shell docker images -f "dangling=true" -q) ; do if [ -n $$i ]; then docker rmi $$i;fi done
	echo -e "\n>>> '编译成功'"
