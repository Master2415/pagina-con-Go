# Utilizar una imagen base de Go oficial
FROM golang:1.20-alpine

# Crear y establecer el directorio de trabajo
WORKDIR /app

# Copiar el archivo go.mod y go.sum si existen
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar el código fuente de la aplicación
COPY . .

# Compilar la aplicación Go
RUN go build -o main .

# Exponer el puerto en el que correrá la aplicación
EXPOSE 8081

# Comando para ejecutar la aplicación
CMD ["./main"]
