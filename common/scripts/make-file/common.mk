SHELL := /bin/bash

COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# 设置ROOT_DIR
ifeq ($(origin ROOT_DIR), undefined)
	ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/../../.. && pwd -P))
endif
