```go run main.go```

If fails on s390x, but it works fine on amd64. 
Remove syscall.O_DIRECT from https://github.com/raulcabello/test-checksum-s390x/blob/main/main.go#L23 to make it work on s390x.
