FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/data-svc cmd/main.go

FROM alpine

COPY --from=0 /app/data-svc /data-svc

COPY --from=0 /app/pkg/config/envs/dev.env /envs/dev.env

EXPOSE 50052

CMD [ "/data-svc" ]