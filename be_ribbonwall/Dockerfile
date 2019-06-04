# build stage
FROM golang:alpine AS build-env
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk add git
# For debugging with netstat
RUN apk add busybox-extras
ADD src /src

# Download dependencies
RUN go get github.com/go-sql-driver/mysql
#RUN go get github.com/aws/aws-sdk-go/aws
#RUN go get github.com/aws/aws-sdk-go/aws/credentials/stscreds
#RUN go get github.com/aws/aws-sdk-go/aws/session
#RUN go get github.com/aws/aws-sdk-go/service/rds/rdsutils

# Build go package
RUN cd /src && go build -o be_ribbonwall main.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/be_ribbonwall /app/

EXPOSE 8080

#ENV variables to load from --build-arg
ARG db_user
ARG db_password
ARG db_name
ARG db_endpoint
ARG aws_region
ARG aws_arn

# DB credentials from ENV
ENV db_user $db_user
ENV db_password $db_password
ENV db_name $db_name
ENV db_endpoint $db_endpoint
ENV aws_region $aws_region
ENV aws_arn $aws_arn

RUN ["chmod", "+x", "./be_ribbonwall"]
ENTRYPOINT ./be_ribbonwall