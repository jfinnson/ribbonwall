# build stage
FROM golang:alpine AS build-env
RUN apk add git

WORKDIR $GOPATH/src/github.com/ribbonwall
COPY . .
RUN go get -d -v ./...
RUN go build -o /app ./be_ribbonwall/main.go

# final stage
FROM alpine
RUN apk --no-cache upgrade && apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build-env /app /app/
RUN mkdir -p /app/be_ribbonwall/config/credentials
COPY ./be_ribbonwall/config/config.production.yaml /app/be_ribbonwall/config
COPY ./be_ribbonwall/config/credentials/ribbonwall.pem /app/be_ribbonwall/config/credentials

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

RUN ["chmod", "+x", "./app"]
ENTRYPOINT ./app