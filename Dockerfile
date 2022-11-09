#FROM golang:1-alpine AS builder
#
#WORKDIR D:/Work/InnowisePreTraineeTask
#
#RUN apk add --update git
#
#COPY cmd/main.go .
#COPY go.mod .
#COPY go.sum .
#COPY internal .
#RUN go mod download
#
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .
#
#FROM alpine
#EXPOSE 8081
#
#COPY --from=builder D:/Work/InnowisePreTraineeTask .
#
#ENTRYPOINT .



FROM golang:alpine AS builder

ADD . /src
RUN cd /src && go build -o main.bin cmd/main.go

# Final step
FROM alpine as runner

WORKDIR /app
COPY --from=builder /src/main.bin /app/

ENTRYPOINT ./main.bin

CMD ["/app"]

#FROM golang:1.18-alpine AS builder
#
#WORKDIR /go/src/your-app
#
#COPY go.mod ./
#COPY go.sum ./
#RUN go mod download
#
#COPY --from=builder /src/main.bin /your-app/
#
#RUN go build -o /your-app
#
#EXPOSE 80
#
#CMD [ "/your-app" ]
