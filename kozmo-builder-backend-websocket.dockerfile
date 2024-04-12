# ---------------------
# build kozmo-builder-backend-websocket
FROM golang:1.20-bullseye as builder-for-backend

## set env
ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

## build
WORKDIR /opt/kozmo/kozmo-builder-backend
RUN cd  /opt/kozmo/kozmo-builder-backend
RUN ls -alh

COPY ./ ./

RUN cat ./Makefile

RUN make build-websocket-server

RUN ls -alh ./bin/kozmo-builder-backend-websocket


# -------------------
# build runner images
FROM alpine:latest as runner

WORKDIR /opt/kozmo/kozmo-builder-backend/bin/

## copy backend bin
COPY --from=builder-for-backend /opt/kozmo/kozmo-builder-backend/bin/kozmo-builder-backend-websocket /opt/kozmo/kozmo-builder-backend/bin/


RUN ls -alh /opt/kozmo/kozmo-builder-backend/bin/



# run
EXPOSE 8002
CMD ["/bin/sh", "-c", "/opt/kozmo/kozmo-builder-backend/bin/kozmo-builder-backend-websocket"]
