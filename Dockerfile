FROM golang

COPY main.go main.go

ENTRYPOINT [ "go", "run", "main.go" ]
