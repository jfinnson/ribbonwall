# build stage
FROM golang:alpine AS build-env
ADD src /src
RUN cd /src && go build -o be_ribbonwall main.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/be_ribbonwall /app/

EXPOSE 8080

RUN ["chmod", "+x", "./be_ribbonwall"]
ENTRYPOINT ./be_ribbonwall