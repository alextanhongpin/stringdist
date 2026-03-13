GIT_COMMIT := $(shell git rev-parse --short HEAD)
start:
	go run cmd/main.go


mem:
	go tool pprof mem.out

cpu:
	go tool pprof cpu.out
