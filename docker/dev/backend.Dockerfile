FROM golang:1.25-alpine
WORKDIR /app

ENV GOTOOLCHAIN=auto

COPY trusmi_be/go.mod trusmi_be/go.sum ./
RUN go mod download

COPY trusmi_be/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/api ./cmd/api

EXPOSE 8090
CMD ["/app/bin/api"]
