FROM golang:1.11 as build
WORKDIR /build
COPY . /build
RUN mkdir /out && \
    CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o /out/echosvc echosvc.go

FROM scratch
COPY --from=build /out/echosvc /
EXPOSE 8080
ENTRYPOINT ["/echosvc"]
