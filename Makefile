# ==============================================================================
# Includes
include common/scripts/make-file/common.mk
include common/scripts/make-file/swagger.mk
include common/scripts/make-file/image.mk

# ==============================================================================

# 创建swagger文档
.PHONY: swagger
swagger:
	@$(MAKE) swagger.run

.PHONY: serve-swagger
serve-swagger:
	@$(MAKE) swagger.server

.PHONY: build
build:
	@$(MAKE) build-image

.PHONY: apply-cm
apply-cm:
	kubectl apply -f deploy/alarmv2-config.yaml

.PHONY: deploy-silence
deploy-silence:
	ssh kunshan-xxxxx "kubectl apply -f deploy/alarmv2.yaml"

.PHONY: deploy-poc
deploy-poc: build
	ssh 10.177.10.1 "kubectl set image deployment/alarmv2 -n xxxxx alarmv2=harbor.inner.galaxy.xxxxx.com/xxxxx/alarmv2:$(VERSION)"