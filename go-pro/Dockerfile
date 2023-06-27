# Usa la imagen base de Go
FROM golang:1.20.5

# Establece el directorio de trabajo
WORKDIR /go-pro

# Copia los archivos go.mod y go.sum
COPY go.mod go.sum ./

# Descarga las dependencias
RUN go mod download

# Copia el resto de los archivos del proyecto
COPY . .

# Agrega manualmente la entrada faltante en el archivo go.sum
RUN echo "github.com/dgrijalva/jwt-go v3.2.0+incompatible" >> go.sum

# Agrega manualmente la entrada faltante en el archivo go.sum
RUN echo "golang.org/x/text v0.3.8" >> go.sum

# Compila el proyecto
RUN CGO_ENABLED=0 go build -o main .

# Expone el puerto 8090
EXPOSE 8090

# Define el comando de inicio
CMD ["./main"]
