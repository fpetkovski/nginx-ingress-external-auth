FROM golang:1.11 as build
WORKDIR /build
COPY . /build
RUN mkdir /out && \
    CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -mod vendor -o /out/authsvc main.go

FROM scratch
COPY --from=build /out/authsvc /
EXPOSE 8080
ENTRYPOINT ["/authsvc"]
