FROM golang:1.25-alpine AS builder
WORKDIR /src

ENV GOTOOLCHAIN=auto

COPY trusmi_be/go.mod trusmi_be/go.sum ./
RUN go mod download
COPY trusmi_be/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o /out/api ./cmd/api

FROM alpine:3.22
WORKDIR /app
RUN apk add --no-cache ca-certificates tzdata \
    && addgroup -S app && adduser -S app -G app
COPY --from=builder /out/api /app/api
USER app
EXPOSE 8090
CMD ["./api"]
