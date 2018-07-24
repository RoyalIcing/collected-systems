FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY . .
RUN apk add --no-cache git
RUN go get -v ./...
# RUN ls /go/src/github.com
# RUN go get github.com/gobuffalo/packr
# RUN ls /go/bin
# RUN /go/bin/packr
RUN go install -v ./...

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
COPY --from=builder /go/src/app/samples /go/src/app/samples

EXPOSE 3838
ENTRYPOINT ./app
