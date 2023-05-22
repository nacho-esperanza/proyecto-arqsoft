# proyecto-arqsoft

### Bajar cambios del proyecto:

```bash
git pull origin main
```


## Comandos de GO

### Corroborar la instalación y versión de GO

```bash
go version
```

### Correr un programa GO

> Pararse siempre en la carpeta que contiene el archivo `main.go`

```bash
go run main.go
```

### Compilar y correr el programa en 2 pasos

> Pararse siempre en la carpeta que contiene el archivo `main.go`<br/>Para Linux o macOS, quitar la extensión `.exe`

```bash
go build
./nombre_del_ejecutable.exe
```

### Crear un módulo para poder usar paquetes

> Pararse siempre en la carpeta que contiene el archivo `main.go`

```bash
go mod init
go mod tidy
```

> Para todos los imports, usar el prefijo `nombre_del_modulo`, por ejemplo `import "go-api/my/package"`