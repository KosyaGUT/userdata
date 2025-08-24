# Используем официальный образ Go
FROM golang:1.25-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы модулей и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код
COPY . .

# Собираем приложение
RUN go build -o main ./cmd/userdata

# Экспонируем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]