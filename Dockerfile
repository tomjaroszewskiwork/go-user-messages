FROM golang:1.10 AS builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/tomjaroszewskiwork/go-user-messages
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /app .

FROM scratch
COPY --from=builder /app ./
ENTRYPOINT ["./app"]