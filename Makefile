gen:
	go generate ./cmd/gen_resources
	go fmt ./generated

arch:
	go run ./cmd/arch/arch.go

dev:
	go build -o /tmp/galaxy-tmp
	/tmp/galaxy-tmp -profile=true -profileport=17999

prof-cpu:
	go tool pprof -http=:8080 /tmp/galaxy-tmp http://0.0.0.0:17999/debug/pprof/profile\?seconds\=3

prof-mem:
	go tool pprof -http=:8080 /tmp/galaxy-tmp http://0.0.0.0:17999/debug/pprof/heap