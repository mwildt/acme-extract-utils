FROM golang:1.23 AS build

RUN mkdir /src
WORKDIR /src
COPY . /src

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o application -tags prod ./cmd/container/main.go
RUN ls -la

FROM scratch
COPY --from=build /src/application /application
ENTRYPOINT ["/application"]