.PHONY: swagger.verify
swagger.verify:
	@if ! which swagger &>/dev/null; then \
		echo "===========> swagger uninstalled"; \
		exit 1; \
	fi

.PHONY: swagger.run
swagger.run: swagger.verify
	@echo "===========> Generateing swagger API docs"
	@swagger generate spec --scan-models -w $(ROOT_DIR) -o $(ROOT_DIR)/swagger/swagger.json
	@swagger generate spec --scan-models -w $(ROOT_DIR) -o $(ROOT_DIR)/swagger/swagger.yaml

.PHONY: swagger.server
swagger.server:
	@swagger serve -F=swagger --no-open --port 36666 $(ROOT_DIR)/swagger/swagger.json