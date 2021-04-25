# Build container
FROM golang:1.16.2-buster AS build_base

RUN apt-get install git

RUN apt install curl &&\
    curl -sL https://deb.nodesource.com/setup_6.x | bash - &&\
    apt-get install -y nodejs &&\
    apt-get install -y npm

WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN cd ./webapp && npm run build

# Unit tests
RUN CGO_ENABLED=0 go test -v

RUN go build -o ./out/scrum-poker .

# Start fresh from a smaller image for the runtime container
FROM debian:buster

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates

RUN update-ca-certificates

WORKDIR /app

USER nobody

COPY --from=build_base --chown=nobody /tmp/app/out/scrum-poker /app/scrum-poker

EXPOSE ${HTTP_PORT}

CMD ["./scrum-poker", "start"]