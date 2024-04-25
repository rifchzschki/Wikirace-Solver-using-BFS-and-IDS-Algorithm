
FROM golang:1.21.6 AS builder


WORKDIR /app


COPY . .


RUN go mod download


# Build aplikasi
RUN go build -o main .


FROM golang:1.21.6

WORKDIR /app



COPY --from=builder /app/main .

COPY Frontend /app/Frontend

EXPOSE 8080

# Command untuk menjalankan aplikasi
CMD ["./main"]
