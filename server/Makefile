DOCKER_USER?=test
IMAGE_NAME?=unshort

.PHONY: generate build dockerize

generate:
	@echo "Go install esc"
	@go get github.com/programmfabrik/esc
	@GOARCH=$(go env GOHOSTARCH) go install github.com/programmfabrik/esc
	@echo "Generating main assets..."
	@esc -private -local-prefix-cwd -pkg=main -o=static.go static/
	@echo "Generating db assets..."
	@cd ./db && esc -private -local-prefix-cwd -pkg=db -o=static.go standard_hosts.txt
	@echo "Assets generated"

build: generate
	@echo "Building..."
	@CGO_ENABLED=1 go build -o unshort.link
	@echo "Build completed. Run the server by ./unshort.link"

test:
	@echo "Running tests...."
	@CGO_ENABLED=1 go test ./...
	@echo "Finished tests"
	@echo "Running vet..."
	@CGO_ENABLED=1 go vet ./...
	@echo "Finished vet"

clean:
	@echo "Started cleaning...."
	@rm static.go blacklist.db link.db unshort.link
	@echo "Finished cleaning"

dockerize:
	@echo "Start dockerizing...."
	docker image build -t $(IMAGE_NAME) .
	@echo "Finished dockerizing"
