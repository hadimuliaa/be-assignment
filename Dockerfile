# Menggunakan gambar dasar golang untuk build
FROM golang:1.20 AS builder

WORKDIR /app

# Menyalin semua file ke dalam kontainer
COPY . .

# Mendownload dependency dan membangun binary aplikasi
RUN go mod tidy
RUN go build -o account-manager main.go

# Debugging: Menampilkan isi direktori /app sebelum tahap final
RUN ls -l /app

# Membuat gambar akhir untuk menjalankan aplikasi
FROM alpine:3.18

WORKDIR /app

# Menginstal dependensi yang mungkin diperlukan, termasuk glibc
RUN apk add --no-cache libc6-compat

# Menyalin binary dari build stage
COPY --from=builder /app/account-manager .
COPY .env .
# Debugging: Menampilkan isi direktori /app setelah tahap final
RUN ls -l /app

# Menjalankan binary
CMD ["./account-manager"]
