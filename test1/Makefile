SAM_TEMPLATE = template.yaml
SAM_CONFIG_FILE = samconfig.toml
STACK_NAME = heartbeat-manager

clean:
	rm -f heartbeat packaged_*.yaml
	rm -rf .aws-sam/

build: clean
	sam build
	cp .aws-sam/build/HeartbeatFunction/heartbeat .

package: build
	sam package \
		--template-file $(SAM_TEMPLATE) \
		--config-file ${SAM_CONFIG_FILE} \
		--config-env ${CONFIG_ENV} \
		--s3-bucket ${S3_BUCKET} \
		--s3-prefix "heartbeat-manager" \
		--region ${REGION} \
		--profile ${PROFILE} \
		--output-template-file packaged_${CONFIG_ENV}.yaml

deploy: build
	sam deploy \
		--confirm-changeset \
		--template-file $(SAM_TEMPLATE) \
		--stack-name ${STACK_NAME} \
		--config-file ${SAM_CONFIG_FILE} \
		--config-env ${CONFIG_ENV} \
		--s3-prefix "heartbeat-manager" \
		--region ${REGION} \
		--profile ${PROFILE} \
		--capabilities CAPABILITY_IAM \
		--force-upload

#deploy:
#	sam deploy \
#		--confirm-changeset \
#		--template-file packaged_${CONFIG_ENV}.yaml \
#		--stack-name ${STACK_NAME} \
#		--region eu-west-1 \
#		--capabilities CAPABILITY_IAM \
#		--profile Dev_AWS_Admin

#deploy:
#	sam deploy \
#		--template-file packaged.yaml \
#		--stack-name ${STACK_NAME} \
#		--config-file ${SAM_CONFIG_FILE} \
#		--config-env development \
#		--region eu-west-1 \
#		--capabilities CAPABILITY_IAM \
#		--profile Dev_AWS_Admin
#
delete: 
	sam delete \
		--stack-name ${STACK_NAME} \
		--config-file ${SAM_CONFIG_FILE} \
		--config-env ${CONFIG_ENV} \
		--s3-prefix "heartbeat-manager" \
		--region ${REGION} \
		--profile ${PROFILE}

#
list:
	sam list resources \
		--stack-name ${STACK_NAME} \
		--config-file ${SAM_CONFIG_FILE} \
		--config-env ${CONFIG_ENV} \
		--region ${REGION} \
		--profile ${PROFILE}

output:
	sam list stack-outputs \
		--stack-name ${STACK_NAME} \
		--config-file ${SAM_CONFIG_FILE} \
		--config-env development \
		--region eu-west-1 \
		--profile ${PROFILE}


dev/deploy: CONFIG_ENV = development
dev/deploy: REGION = eu-west-1
dev/deploy: S3_BUCKET = aws-sam-cli-managed-default-samclisourcebucket-qcylqszmdjzp
dev/deploy: PROFILE = Dev_AWS_Admin
dev/deploy: deploy

dev/delete: CONFIG_ENV = development
dev/delete: REGION = eu-west-1
dev/delete: S3_BUCKET = aws-sam-cli-managed-default-samclisourcebucket-qcylqszmdjzp
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

