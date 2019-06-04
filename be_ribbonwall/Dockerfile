# build stage
FROM golang:alpine AS build-env
#RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk --no-cache upgrade && apk --no-cache add ca-certificates
RUN apk add git
ADD src /src


# add AWS RDS CA bundle
#ADD https://s3.amazonaws.com/rds-downloads/rds-combined-ca-bundle.pem /tmp/rds-ca/aws-rds-ca-bundle.pem
## split the bundle into individual certs (prefixed with xx)
## see http://blog.swwomm.com/2015/02/importing-new-rds-ca-certificate-into.html
#RUN cd /tmp/rds-ca && cat aws-rds-ca-bundle.pem|awk 'split_after==1{n++;split_after=0} /-----END CERTIFICATE-----/ {split_after=1} {print > "cert" n ""}' \
#    && for CERT in /tmp/rds-ca/cert*; do mv $CERT /usr/local/share/ca-certificates/aws-rds-ca-$(basename $CERT).crt; done \
#    && rm -rf /tmp/rds-ca \
#&& update-ca-certificates

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