
# BUILD NUXT TO SERVABLE
FROM node:20 AS nuxt-builder

WORKDIR /builder


COPY ./package*.json ./
COPY ./yarn.lock yarn.lock

# RUN npm i
RUN yarn

ADD ./app ./app

RUN yarn build:nuxt

# BUILD POCKTBASE WITH SERVABLE TO BINARY
FROM golang:1.21.2 AS go-builder

RUN apt-get update && apt-get install -y ca-certificates openssl

ARG cert_location=/usr/local/share/ca-certificates
ARG crt_domain="api.telegram.org"
RUN openssl s_client -showcerts -connect ${crt_domain}:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/directus.crt
RUN update-ca-certificates

WORKDIR /builder
COPY go.mod go.sum ./
RUN go mod download

ADD ./pocketbase ./pocketbase

COPY --from=nuxt-builder ./builder/app/.output/public ./pocketbase/public
# COPY --from=nuxt-builder ./builder/app/.output/public/assets ./pocketbase/assets

RUN go build -o pocketnuxt ./pocketbase/.

# SERVE PROJECT
FROM debian:latest
RUN apt update && apt upgrade -y 
# && apt install -y ca-certificates iptables && rm -rf /var/lib/apt/lists/*
# RUN iptables -A INPUT -p tcp --dport 80 -j ACCEPT

COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /
COPY --from=go-builder /builder/pocketnuxt /pocketnuxt
RUN chmod +x /pocketnuxt
EXPOSE 80
# USER nonroot:nonroot
CMD ["/pocketnuxt","serve","--http","0.0.0.0:80"]
