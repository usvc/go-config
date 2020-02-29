PROJECT_NAME=config
CMD_ROOT=config
DOCKER_NAMESPACE=usvc
DOCKER_IMAGE_NAME=libeg-config

-include ./makefile.properties

deps:
	go mod vendor -v
	go mod tidy -v
run:
	go run ./cmd/$(CMD_ROOT)
test:
	go test -v ./... -cover -coverprofile c.out
build:
	CGO_ENABLED=0 \
	go build \
		-v \
		-o ./bin/$(CMD_ROOT)_$$(go env GOOS)_$$(go env GOARCH)${BIN_EXT} \
		./cmd/$(CMD_ROOT)
build_production:
	CGO_ENABLED=0 \
	go build \
		-a -v \
		-ldflags "-X main.Commit=$$(git rev-parse --verify HEAD) \
			-X main.Version=$$(git describe --tag $$(git rev-list --tags --max-count=1)) \
			-X main.Timestamp=$$(date +'%Y%m%d%H%M%S') \
			-extldflags 'static' \
			-s -w" \
		-o ./bin/$(CMD_ROOT)_$$(go env GOOS)_$$(go env GOARCH)${BIN_EXT} \
		./cmd/$(CMD_ROOT)
compress_production:
	upx -9 -v ./bin/$(CMD_ROOT)_$$(go env GOOS)_$$(go env GOARCH)${BIN_EXT}
	upx -t ./bin/$(CMD_ROOT)_$$(go env GOOS)_$$(go env GOARCH)${BIN_EXT}

image:
	docker build --file ./deploy/Dockerfile --tag $(DOCKER_NAMESPACE)/$(DOCKER_IMAGE_NAME):latest .
save:
	mkdir -p ./build
	docker save --output ./build/$(PROJECT_NAME).tar.gz $(DOCKER_NAMESPACE)/$(DOCKER_IMAGE_NAME):latest
load:
	docker load --input ./build/$(PROJECT_NAME).tar.gz
dockerhub:
	docker push $(DOCKER_NAMESPACE)/$(DOCKER_IMAGE_NAME):latest
	git fetch
	docker tag $(DOCKER_NAMESPACE)/$(DOCKER_IMAGE_NAME):latest \
		$(DOCKER_NAMESPACE)/$(DOCKER_IMAGE_NAME):$$(git describe --tag $$(git rev-list --tags --max-count=1))
	docker push $(DOCKER_NAMESPACE)/$(DOCKER_IMAGE_NAME):$$(git describe --tag $$(git rev-list --tags --max-count=1))

see_ci:
	xdg-open https://gitlab.com/usvc/modules/go/config/pipelines

.ssh:
	mkdir -p ./.ssh
	ssh-keygen -t rsa -b 8192 -f ./.ssh/id_rsa -q -N ""
	cat ./.ssh/id_rsa | base64 -w 0 > ./.ssh/id_rsa.base64
