# run with:
# docker build -t malasgit .
# docker run -it malasgit:latest /bin/sh

FROM golang:1.25 as build
WORKDIR /go/src/github.com/itokun99/malasgit/
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine:3.19
RUN apk add --no-cache -U git xdg-utils
WORKDIR /go/src/github.com/itokun99/malasgit/
COPY --from=build /go/src/github.com/itokun99/malasgit ./
COPY --from=build /go/src/github.com/itokun99/malasgit/malasgit /bin/
RUN echo "alias gg=malasgit" >> ~/.profile

ENTRYPOINT [ "malasgit" ]
