GIT_COMMIT := $(shell git rev-parse --short HEAD)
start:
	go run cmd/main.go -cpu "./profiling/cpu_${GIT_COMMIT}.out" -mem "./profiling/mem_${GIT_COMMIT}.out"
	# go run $(shell ls -1 *.go | grep -v _test.go) -cpu "./profiling/cpu_${GIT_COMMIT}.out" -mem "./profiling/mem_${GIT_COMMIT}.out"
