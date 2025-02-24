FROM golang:1.24

WORKDIR /app/backend

# Instalar Air para hot-reload (nueva ruta)
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE $PORT

# Usar Air para desarrollo
CMD ["air"]
