FROM golang:latest
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
# ENV PORT 8080
RUN go build
EXPOSE 8080
# CMD ["go run ."]
