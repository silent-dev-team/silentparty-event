
# BUILD NUXT TO SERVABLE
FROM node:20 as nuxt-builder

WORKDIR /builder

COPY ./package*.json ./
COPY ./yarn.lock yarn.lock

RUN npm i
RUN yarn

ADD ./app ./app

RUN yarn build:nuxt

# BUILD POCKTBASE WITH SERVABLE TO BINARY
FROM golang:1.21.2 as go-builder

WORKDIR /builder
COPY go.mod go.sum ./
RUN go mod download

ADD ./pocketbase ./pocketbase

COPY --from=nuxt-builder ./builder/app/.output/public ./pocketbase
RUN go build -o pocketnuxt ./pocketbase/.

# SERVE PROJECT
FROM debian:latest as live-system
WORKDIR /
COPY --from=go-builder /builder/pocketnuxt /pocketnuxt
EXPOSE 8090
#USER nonroot:nonroot
CMD ["/pocketnuxt","serve","--http","127.0.0.1:8090"]