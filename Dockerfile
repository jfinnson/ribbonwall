# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache upgrade && apk --no-cache add ca-certificates
RUN apk add git
ADD be_ribbonwall /be_ribbonwall
ADD common /common


# Download dependencies
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gin-contrib/cors
RUN go get github.com/gin-contrib/sessions
RUN go get github.com/gin-contrib/sessions/cookie
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/jinzhu/gorm
RUN go get github.com/jinzhu/configor
RUN go get github.com/satori/go.uuid
RUN go get github.com/sirupsen/logrus
RUN go get github.com/auth0-community/go-auth0
RUN go get gopkg.in/square/go-jose.v2

# Build go package
#RUN cd /be_ribbonwall && go build
RUN go build ./be_ribbonwall/...

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /be_ribbonwall /app/

EXPOSE 8080

#ENV variables to load from --build-arg
ARG db_user
ARG db_password
ARG db_name
ARG db_host
ARG db_port
ARG aws_region
ARG aws_arn
ARG auth_client_secret

# Genetal ENVs
ENV SERVICE_CONFIG production
# DB credentials from ENV
ENV DB_USER $db_user
ENV DB_PASSWORD $db_password
ENV DB_NAME $db_name
ENV DB_HOST $db_host
ENV DB_PORT $db_port
ENV AWS_REGION $aws_region
ENV AWS_ARN $aws_arn
# AUTH credentials from ENV
ENV AUTH_CLIENT_SECRET $auth_client_secret

RUN ["chmod", "+x", "./be_ribbonwall"]
ENTRYPOINT ./be_ribbonwall