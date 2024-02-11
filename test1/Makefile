SAM_TEMPLATE = template.yaml

.PHONY: build
build:
	sam build

deploy_dev: build
	sam deploy \
		--template-file packaged.yaml \
		--config-file samconfig.toml \
		--config-env development \
		--region eu-west-1 \
		--capabilities CAPABILITY_IAM \
		--profile Dev_AWS_Admin

#invoke:
#	sam local invoke $(FUNCTION_NAME) \
#		--event event.json

#logs:
#	sam logs -n $(FUNCTION_NAME)
#
#test:
#	sam local invoke $(FUNCTION_NAME) \
#		--event event.json

# Define phony targets
.PHONY: build package deploy invoke logs test

