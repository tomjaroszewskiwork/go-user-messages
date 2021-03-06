# Go User Messages

This is a go version of [user-messages](https://github.com/tomjaroszewskiwork/user-messages).

First go-lang app! Don't be too mean.

## To build

Get Go 1.10
Get [dep](https://github.com/golang/dep)

```bash
dep ensure
go build
```

Because persistance is provided by sqllite gcc is required on PATH.

## To test

```bash
go test ./...
```

## To build using docker

```bash
docker build -t go-user-messages .
```

## To run docker image

```bash
docker run -p 8080:8080 go-user-messages
```

## API spec

Default port is 8080.

Sample API call: http://localhost:8080/v1/users/tom.j/messages

Please see [swagger.yaml](https://github.com/tomjaroszewskiwork/go-user-messages/blob/master/swagger.yaml) for full API spec details.

## Architecture Overview

![Architecture](/architecture.svg)

Application runs Gorrilla mux router listening in on port 8080.

REST API method handlers then call on store package which wraps the calls around GORM.

Finally the data is stored in a sql lite database.
