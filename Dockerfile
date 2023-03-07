# Start from a Golang base image
FROM golang:1.18.10-alpine3.17 AS build
WORKDIR /application
COPY ./ ./
RUN go mod tidy
RUN mkdir -p /opt
RUN go build -o /opt/application cmd/main.go
# Start from a new base image without the Golang tools
FROM alpine:3.17 AS runtime
# Copy the binary from the previous stage
COPY --from=build /opt/application /usr/local/bin/versioner
RUN chmod a+x /usr/local/bin/versioner
ENTRYPOINT ["versioner"]
CMD ["help"]