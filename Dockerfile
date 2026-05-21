FROM golang:1.25-alpine AS build

ARG VERSION=0.0.0-local

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X main.version=${VERSION}" -o /app/backend .

FROM scratch

COPY --from=build /app/backend /backend

EXPOSE 8080

ENTRYPOINT ["/backend"]
