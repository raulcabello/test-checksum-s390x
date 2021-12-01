go run main.go

If fails on s390x, but it works fine on amd64. 
Remove syscall.O_DIRECT to make it work on s390x.
