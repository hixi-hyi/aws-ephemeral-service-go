.PHONY: localstack
localstack: clean
	docker-compose -f localstack/docker-compose.yaml up
tests:
	AWS_DEFAULT_REGION=ap-northeast-1 AWS_REGION=ap-northeast-1 AWS_PROFILE=localstack go test ./awseph/... -v -count=1
clean:
	rm -rf ./localstack/tmp
