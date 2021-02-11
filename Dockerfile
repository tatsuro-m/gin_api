FROM golang:1.15.7-alpine as dev

ENV ROOT=/app
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download
COPY . ${ROOT}

CMD ["go", "run", "main.go"]


FROM golang:1.15.7-alpine as builder

ENV ROOT=/app
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download

COPY . ${ROOT}
RUN CGO_ENABLED=0 GOOS=linux go build -o $ROOT/binary


FROM scratch as prod

ENV ROOT=/app
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/binary ${ROOT}

CMD ["/app/binary"]
