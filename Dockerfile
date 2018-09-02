FROM golang:1.10

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]


# # STEP 1 build executable binary

# FROM golang:alpine as builder
# COPY . $GOPATH/src/github.com/tomjaroszewskiwork/go-user-messages
# WORKDIR $GOPATH/src/github.com/tomjaroszewskiwork/go-user-messages

# RUN apk add --no-cache gcc musl-dev
# RUN apk add --no-cache git

# #get dependancies
# #you can also use dep
# RUN go get -d -v

# #build the binary
# RUN go build -o /go/bin/hello

# # STEP 2 build a small image

# # start from scratch
# FROM scratch

# # Copy our static executable
# COPY --from=builder /go/bin/hello /go/bin/hello
# ENTRYPOINT ["/go/bin/hello"]