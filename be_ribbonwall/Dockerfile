# build stage
FROM golang:alpine AS build-env
ADD src /src
RUN cd /src && go build -o be_ribbonwall main.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/be_ribbonwall /app/

EXPOSE 8080

# DB credentials from ENV
ENV db_name $db_user
ENV db_name $db_name
ENV db_endpoint $db_endpoint
ENV aws_region $aws_region
ENV aws_arn $aws_arn

RUN ["chmod", "+x", "./be_ribbonwall"]
ENTRYPOINT ./be_ribbonwall