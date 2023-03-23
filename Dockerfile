from golang as builder
WORKDIR /go/src/github.com/mgrusso/ubi-hello
COPY ./main.go ./go.mod ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM registry.access.redhat.com/ubi8/ubi-micro
WORKDIR / 
COPY --from=builder /go/src/github.com/mgrusso/ubi-hello/app /
EXPOSE 8080
CMD ["./app"]