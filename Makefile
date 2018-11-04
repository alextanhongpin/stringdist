GIT_COMMIT := $(shell git rev-parse --short HEAD)
start:
	go run $(shell ls -1 *.go | grep -v _test.go) -cpu "cpu_${GIT_COMMIT}.out" -mem "mem_${GIT_COMMIT}.out"
