.PHONY: localstack
localstack:
	docker-compose -f localstack/docker-compose.yaml up
tests:
	AWS_DEFAULT_REGION=ap-northeast-1 AWS_REGION=ap-northeast-1 AWS_PROFILE=localstack go test ./awsephemeral/... -v -count=1
