SAM_TEMPLATE = template.yaml
STACK_NAME = heartbeat-manager

clean:
	rm -rf .aws-sam/

build: clean
	sam build

deploy: build
	sam deploy \
		--s3-prefix heartbeat-manager \
		--region ${REGION} \
		--capabilities CAPABILITY_IAM \
		--profile ${PROFILE}

delete: 
	sam delete \
		--s3-prefix heartbeat-manager \
		--region ${REGION} \
		--profile ${PROFILE}

#
list:
	sam list resources \
		--region ${REGION} \
		--profile ${PROFILE}

output:
	sam list stack-outputs \
		--region eu-west-1 \
		--profile ${PROFILE}


dev/deploy: CONFIG_ENV = development
dev/deploy: REGION = eu-west-1
dev/deploy: PROFILE = Dev_AWS_Admin
dev/deploy: deploy

dev/delete: CONFIG_ENV = development
dev/delete: REGION = eu-west-1
dev/delete: PROFILE = Dev_AWS_Root
dev/delete: delete

dev/list: CONFIG_ENV = development
dev/list: REGION = eu-west-1
dev/list: PROFILE = Dev_AWS_Admin
dev/list: list

dev/output: REGION = eu-west-1
dev/output: CONFIG_ENV = development
dev/output: PROFILE = Dev_AWS_Admin
dev/output: output

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
.PHONY: build package deploy invoke logs test clean install update lambda hearbeat
.PHONY: dev/deploy dev/list dev/output

