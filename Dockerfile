
# BUILD NUXT TO SERVABLE
FROM node:20 as nuxt-builder

WORKDIR /builder


COPY ./package*.json ./
COPY ./yarn.lock yarn.lock

RUN npm i
RUN yarn

ADD ./app ./app

ENV NUXT_POCKETBASE_URL "https://app.silentparty-hannover.de"
RUN yarn build:nuxt

# BUILD POCKTBASE WITH SERVABLE TO BINARY
FROM golang:1.21.2 as go-builder

WORKDIR /builder
COPY go.mod go.sum ./
RUN go mod download

ADD ./pocketbase ./pocketbase

COPY --from=nuxt-builder ./builder/app/.output/public ./pocketbase/public
RUN go build -o pocketnuxt ./pocketbase/.

# SERVE PROJECT
FROM debian:latest
RUN apt update && apt upgrade -y 
# && apt install -y ca-certificates iptables && rm -rf /var/lib/apt/lists/*
# RUN iptables -A INPUT -p tcp --dport 80 -j ACCEPT
WORKDIR /
COPY --from=go-builder /builder/pocketnuxt /pocketnuxt
RUN chmod +x /pocketnuxt
EXPOSE 80
# USER nonroot:nonroot
CMD ["/pocketnuxt","serve","--http","0.0.0.0:80"]
