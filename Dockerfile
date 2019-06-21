# build stage
#FROM golang:1.12 AS build-env
FROM golang:1.12 AS build-env
#FROM golang:alpine AS build-env
#RUN apk add git
#RUN apt-get update

# Install npm
RUN curl -sL https://deb.nodesource.com/setup_12.x | bash -
RUN apt-get install -y nodejs
RUN npm install gulp -g
RUN npm install yarn -g
RUN apt-get install git

WORKDIR $GOPATH/src/github.com/jfinnson/ribbonwall


## be_ribbonwall
# Force the go compiler to use modules
ENV GO111MODULE=on
## We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
#This is the ‘magic’ step that will download all the dependencies that are specified in
# the go.mod and go.sum file.
# Because of how the layer caching system works in Docker, the  go mod download
# command will _ only_ be re-run when the go.mod or go.sum file change
# (or when we add another docker instruction this line)
RUN go mod download
## Build be_ribbonwalld
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app ./domains/be_ribbonwall/main.go



# fe_competitors
RUN cd ./domains/fe_competitors && npm install && npm run build
# fe_admin
RUN cd ./domains/fe_admin && npm install && npm run build
#&& rm build/static/js/runtime*
#RUN rsync -a ./domains/fe_admin/build/static/ ./domains/fe_competitors/build/static/

# final stage
FROM alpine
RUN apk --no-cache upgrade && apk --no-cache add ca-certificates
WORKDIR /app
# be_ribbonwall
COPY --from=build-env /app .
RUN mkdir -p /app/domains/be_ribbonwall/config/credentials
COPY ./domains/be_ribbonwall/config/config.production.yaml /app/domains/be_ribbonwall/config
COPY ./domains/be_ribbonwall/config/credentials/ribbonwall.pem /app/domains/be_ribbonwall/config/credentials
# fe_competitors
RUN mkdir -p /app/domains/fe_competitors/build
COPY --from=build-env /go/src/github.com/jfinnson/ribbonwall/domains/fe_competitors/build /app/domains/fe_competitors/build
# fe_admin
RUN mkdir -p /app/domains/fe_admin/build
COPY --from=build-env /go/src/github.com/jfinnson/ribbonwall/domains/fe_admin/build /app/domains/fe_admin/build

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

#RUN ["chmod", "+x", "./app"]
#ENTRYPOINT ./app
CMD ["./app"]