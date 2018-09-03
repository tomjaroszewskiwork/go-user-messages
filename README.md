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

go test ./...

## To build using docker



## To run docker image



## API spec

Default port is 8080.

Sample API call: http://localhost:8080/v1/users/tom.j/messages

Please see swagger.json for full API spec details.

## Architecture Overview

![Architecture](/architecture.svg)

Application runs Gorrilla mux router listening in on port 8080.

REST API method handlers then call on store package which wraps the calls around GORM.

Finally the data is stored in a sql lite database.

When deployed into Elastic Beanstalk in AWS there is a ngix load balancer in front of the application.









