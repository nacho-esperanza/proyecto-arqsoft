# Usa la imagen base de Go
FROM golang:1.16.5-alpine3.13

# Establece el directorio de trabajo
WORKDIR /go-pro

# Copia los archivos go.mod y go.sum
COPY go.mod go.sum ./

# Descarga las dependencias
RUN go mod download

# Copia el resto de los archivos del proyecto
COPY . .

# Compila el proyecto
RUN CGO_ENABLED=0 go build -o main .

# Expone el puerto 8090
EXPOSE 8090

# Define el comando de inicio
CMD ["./main"]
