# build stage
FROM golang:alpine AS build-env
RUN apk add git

WORKDIR $GOPATH/src/github.com/jfinnson/ribbonwall

## be_ribbonwall
# Force the go compiler to use modules
ENV GO111MODULE=on

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
#This is the ‘magic’ step that will download all the dependencies that are specified in
# the go.mod and go.sum file.
# Because of how the layer caching system works in Docker, the  go mod download
# command will _ only_ be re-run when the go.mod or go.sum file change
# (or when we add another docker instruction this line)
RUN go mod download
# Build be_ribbonwalld
COPY . .
RUN go build -o /app ./domains/be_ribbonwall/main.go

# fe_competitors

# final stage
FROM alpine
RUN apk --no-cache upgrade && apk --no-cache add ca-certificates
WORKDIR /app
# be_ribbonwall
COPY --from=build-env /app /app/
RUN mkdir -p /app/domains/be_ribbonwall/config/credentials
COPY ./domains/be_ribbonwall/config/config.production.yaml /app/domains/be_ribbonwall/config
COPY ./domains/be_ribbonwall/config/credentials/ribbonwall.pem /app/domains/be_ribbonwall/config/credentials
# fe_competitors
#RUN mkdir -p /app/domains/fe_competitors/build
#COPY --from=build-env ./domains/fe_competitors/build /app/domains/fe_competitors/build

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
ENV SERVICE_CONFIG local
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