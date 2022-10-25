FROM golang:1.16-alpine as builder

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o balance-app ./cmd/api/

FROM scratch

COPY --from=builder /build/balance-app /app/
