gen:
	go generate ./cmd/gen_resources
	go fmt ./generated

test:
	go test ./...

lint:
	golangci-lint run

arch:
	docker run --rm \
		-v ${PWD}:/app \
		fe3dback/go-arch-lint:latest-stable-release check --project-path /app

quality-check:
	@echo "=======[ ARCH ] ========"
	make arch
	@echo "=======[ TEST ] ========"
	make test
	@echo "=======[ LINT ] ========"
	make lint
	@echo "                            "
	@echo " ~~~~~~~~~~~~~~~~~~~~~~~~~~ "
	@echo "          All tests passed! "
	@echo " ~~~~~~~~~~~~~~~~~~~~~~~~~~ "

dev:
	go build -o /tmp/galaxy-tmp
	/tmp/galaxy-tmp -profile=true -profileport=17999

prof-cpu:
	go tool pprof -http=:8080 /tmp/galaxy-tmp http://0.0.0.0:17999/debug/pprof/profile\?seconds\=3

prof-mem:
	go tool pprof -http=:8080 /tmp/galaxy-tmp http://0.0.0.0:17999/debug/pprof/heap