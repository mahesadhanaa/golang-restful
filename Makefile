.PHONY: ci-docker-image
ci-docker-image:
	docker build -t "$(CI_DOCKER_IMAGE_OUTPUT)" -f ./.ci/Dockerfile .

run:
	go run main.go

# test:
# 	go test ./...